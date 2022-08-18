package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/victorbetoni/moonitora/authorization"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/repository"
	"net/http"
	"strings"
	"time"
)

func PostMonitoria(c *gin.Context) (int, error) {
	var monitoria model.Monitoria
	if err := c.BindJSON(&monitoria); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	var horario model.Horario
	if err := repository.DownloadHorario(monitoria.Horario, &horario); err != nil {
		return http.StatusInternalServerError, errors.New(err.Error())
	}

	var monitor model.Monitor
	if err := repository.DownloadMonitor(monitoria.Monitor, &monitor); err != nil {
		return http.StatusInternalServerError, errors.New(err.Error())
	}

	marcaPorEmail, _ := authorization.ExtractUser(c)
	monitoria.MarcadaPor = marcaPorEmail

	date, _ := time.Parse("2006-01-02", monitoria.Data)
	if date.Before(time.Now()) {
		return http.StatusBadRequest, errors.New("Essa data já passou.")
	}

	if monitoria.Departamento != monitor.Departamento {
		return http.StatusBadRequest, errors.New("departamento não corresponde")
	}

	if int(date.Weekday()) != horario.DiaDaSemana {
		return http.StatusBadRequest, errors.New("dia da semana não corresponde")
	}

	monitoria.Id = strings.ReplaceAll(uuid.New().String(), "-", "")[:10]
	monitoria.Status = 0

	if !repository.CheckDisponibility(monitoria.Horario, monitoria.Data) {
		return http.StatusBadRequest, errors.New("Dia e horario do monitor ja ocupado")
	}

	if err := repository.InsertMonitoria(monitoria); err != nil {
		return http.StatusInternalServerError, err
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Monitoria marcada com sucesso!", "body": monitoria})
	return 0, nil
}

func FetchMonitorias(c *gin.Context) (int, error) {
	monitor, ok := c.GetQuery("monitor")
	if !ok {
		return http.StatusBadRequest, errors.New("Especifique um monitor")
	}

	type MonitoriaComHorario struct {
		Monitoria model.Monitoria `json:"monitoria"`
		Horario   model.Horario   `json:"horario"`
	}

	var monitoriasComHorario []MonitoriaComHorario

	var monitorias []model.Monitoria
	if err := repository.DownloadMonitorias(monitor, &monitorias); err != nil {
		return http.StatusInternalServerError, err
	}

	for _, monitoria := range monitorias {
		var horario model.Horario
		if err := repository.DownloadHorario(monitoria.Horario, &horario); err != nil {
			return http.StatusInternalServerError, err
		}
		monitoriasComHorario = append(monitoriasComHorario, MonitoriaComHorario{Monitoria: monitoria, Horario: horario})
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "", "body": monitoriasComHorario})
	return 0, nil
}

func CheckDisponibility(c *gin.Context) (int, error) {
	type Request struct {
		Data    string `json:"data"`
		Horario string `json:"horario"`
	}

	var req Request
	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	if !repository.CheckDisponibility(req.Horario, req.Data) {
		return http.StatusBadRequest, errors.New("Dia e horario do monitor ja ocupado")
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "", "body": ""})
	return 0, nil
}

func ConfirmMonitoria(c *gin.Context) (int, error) {
	value, ok := c.GetQuery("monitoria")
	if !ok {
		return http.StatusBadRequest, errors.New("Informe uma monitoria para confirmar")
	}

	user, _ := authorization.ExtractUser(c)

	var sender model.Monitor
	if err := repository.DownloadMonitor(user, &sender); err != nil {
		return http.StatusInternalServerError, err
	}

	var monitoria model.Monitoria
	if err := repository.DownloadMonitoria(value, &monitoria); err != nil {
		return http.StatusInternalServerError, err
	}

	if user != monitoria.Monitor && sender.Adm != 1 {
		return http.StatusUnauthorized, errors.New("Você não tem permissão para isso")
	}

	switch monitoria.Status {
	case 1:
		return http.StatusConflict, errors.New("Essa monitoria já esta confirmada")
	case 2:
		return http.StatusConflict, errors.New("Essa monitoria já foi concluida")
	case 3:
		return http.StatusConflict, errors.New("Essa monitoria já foi cancelada")
	}

	if err := repository.SetStatus(monitoria.Id, 1); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Monitoria confirmada com sucesso", "body": ""})
	return 0, nil
}

func ConcludeMonitoria(c *gin.Context) (int, error) {
	value, ok := c.GetQuery("monitoria")
	if !ok {
		return http.StatusBadRequest, errors.New("Informe uma monitoria para confirmar")
	}

	user, _ := authorization.ExtractUser(c)

	var sender model.Monitor
	if err := repository.DownloadMonitor(user, &sender); err != nil {
		return http.StatusInternalServerError, err
	}

	var monitoria model.Monitoria
	if err := repository.DownloadMonitoria(value, &monitoria); err != nil {
		return http.StatusInternalServerError, err
	}

	if user != monitoria.Monitor && sender.Adm != 1 {
		return http.StatusUnauthorized, errors.New("Você não tem permissão para isso")
	}

	switch monitoria.Status {
	case 1:
		return http.StatusConflict, errors.New("Essa monitoria já esta confirmada")
	case 2:
		return http.StatusConflict, errors.New("Essa monitoria já foi concluida")
	case 3:
		return http.StatusConflict, errors.New("Essa monitoria já foi cancelada")
	}

	if err := repository.SetStatus(monitoria.Id, 1); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Monitoria concluida com sucesso", "body": ""})
	return 0, nil
}

func CancelMonitoria(c *gin.Context) (int, error) {
	value, ok := c.GetQuery("monitoria")
	if !ok {
		return http.StatusBadRequest, errors.New("Informe uma monitoria para confirmar")
	}

	user, _ := authorization.ExtractUser(c)

	var sender model.Monitor
	if err := repository.DownloadMonitor(user, &sender); err != nil {
		return http.StatusInternalServerError, err
	}

	var monitoria model.Monitoria
	if err := repository.DownloadMonitoria(value, &monitoria); err != nil {
		return http.StatusInternalServerError, err
	}

	if user != monitoria.Monitor && sender.Adm != 1 {
		return http.StatusUnauthorized, errors.New("Você não tem permissão para isso")
	}

	switch monitoria.Status {
	case 2:
		return http.StatusConflict, errors.New("Essa monitoria já foi concluida")
	case 3:
		return http.StatusConflict, errors.New("Essa monitoria já foi cancelada")
	}

	if err := repository.SetStatus(monitoria.Id, 3); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "A monitoria foi cancelada e o aluno foi avisado por email.", "body": ""})
	return 0, nil
}
