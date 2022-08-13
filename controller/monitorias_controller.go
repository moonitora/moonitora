package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/repository"
	"net/http"
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

	date, _ := time.Parse("2006-01-02", monitoria.Data)

	if monitoria.Departamento != monitor.Departamento {
		return http.StatusBadRequest, errors.New("departamento não corresponde")
	}

	if int(date.Weekday()) != horario.DiaDaSemana {
		return http.StatusBadRequest, errors.New("dia da semana não corresponde")
	}

	return 0, nil
}
