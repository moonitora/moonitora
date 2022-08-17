package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/victorbetoni/moonitora/authorization"
	"github.com/victorbetoni/moonitora/database"
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
	fmt.Println("1")

	var horario model.Horario
	if err := repository.DownloadHorario(monitoria.Horario, &horario); err != nil {
		return http.StatusInternalServerError, errors.New(err.Error())
	}
	fmt.Println("2")

	var monitor model.Monitor
	if err := repository.DownloadMonitor(monitoria.Monitor, &monitor); err != nil {
		return http.StatusInternalServerError, errors.New(err.Error())
	}
	fmt.Println("3")

	marcaPorEmail, _ := authorization.ExtractUser(c)
	monitoria.MarcadaPor = marcaPorEmail

	date, _ := time.Parse("2006-01-02", monitoria.Data)
	if date.Before(time.Now()) {
		return http.StatusBadRequest, errors.New("Essa data já passou.")
	}

	fmt.Println("4")

	if monitoria.Departamento != monitor.Departamento {
		return http.StatusBadRequest, errors.New("departamento não corresponde")
	}
	fmt.Println("5")

	if int(date.Weekday()) != horario.DiaDaSemana {
		return http.StatusBadRequest, errors.New("dia da semana não corresponde")
	}

	monitoria.Id = strings.ReplaceAll(uuid.New().String(), "-", "")[:10]
	monitoria.Status = 0

	fmt.Println("seila")

	db := database.GrabDB()
	var foo model.Monitoria
	if err := db.Get(&foo, "SELECT * FROM monitorias WHERE horario=$1 AND monitor=$2 AND data=$3", monitoria.Horario, monitoria.Monitor, monitoria.Data); err == nil || (err != nil && err != sql.ErrNoRows) {
		return http.StatusBadRequest, errors.New("Dia e horario do monitor ja ocupado")
	}

	fmt.Println("6")

	if err := repository.InsertMonitoria(monitoria); err != nil {
		return http.StatusInternalServerError, err
	}
	fmt.Println("7")

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
		Day     string `json:"dia"`
		Horario string `json:"horario"`
	}

	var req Request
	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	db := database.GrabDB()
	if err := db.Get(&model.Monitoria{}, "SELECT * FROM monitorias WHERE data=$1 AND horario=$2", req.Day, req.Horario); err == nil {
		return http.StatusBadRequest, errors.New("Dia e horários já reservados")
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "", "body": ""})
	return 0, nil
}
