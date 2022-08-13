package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/repository"
	"net/http"
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
