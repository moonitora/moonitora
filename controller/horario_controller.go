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
)

func FetchHorarios(c *gin.Context) (int, error) {
	monitor, ok := c.GetQuery("monitor")
	if !ok {
		return http.StatusBadRequest, errors.New("especifique um monitor")
	}

	var horarios []model.Horario
	if err := repository.DownloadHorarios(monitor, &horarios); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, horarios)
	return 0, nil
}

func PostHorario(c *gin.Context) (int, error) {
	var horario model.Horario
	if err := c.BindJSON(&horario); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	if !(horario.InicioHoras >= 0 && horario.InicioHoras <= 23 && horario.InicioMinutos >= 0 && horario.InicioMinutos <= 59 && horario.TerminoHoras >= 0 && horario.TerminoHoras <= 23 && horario.TerminoMinutos >= 0 && horario.TerminoMinutos <= 59 && horario.DiaDaSemana >= 0 && horario.DiaDaSemana <= 6) {
		return http.StatusBadRequest, errors.New("bad request")
	}

	user, _ := authorization.ExtractUser(c)
	var sender model.Monitor
	if err := repository.DownloadMonitor(user, &sender); err != nil {
		return http.StatusInternalServerError, err
	}

	if horario.Monitor != user && sender.Adm == 1 {
		return http.StatusUnauthorized, errors.New("Você não tem permissão para isso")
	}

	horario.Id = strings.ReplaceAll(uuid.New().String(), "-", "")[:10]

	if err := repository.InsertHorario(horario); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, horario)
	return 0, nil
}
