package controller

import (
	"errors"
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

	if monitoria.Departamento != monitor.Departamento {
		return http.StatusBadRequest, errors.New("departamento não corresponde")
	}

	if int(date.Weekday()) != horario.DiaDaSemana {
		return http.StatusBadRequest, errors.New("dia da semana não corresponde")
	}

	monitoria.Id = strings.ReplaceAll(uuid.New().String(), "-", "")[:10]

	db := database.GrabDB()
	if err := db.Get(&model.Monitoria{}, "SELECT * FROM monitorias WHERE horario=$1 AND monitor=$2 AND data=$3"); err == nil {
		return http.StatusConflict, errors.New("dia e horario do monitor ja ocupado")
	}

	if err := repository.InsertMonitoria(monitoria); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Monitoria marcada com sucesso!", "body": monitoria})
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
		return http.StatusBadRequest, errors.New("Horário já reservado")
	}

	return 0, nil
}
