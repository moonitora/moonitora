package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/repository"
	"net/http"
)

func FetchHorarios(c *gin.Context) error {
	monitor, ok := c.GetQuery("monitor")
	if !ok {
		return errors.New("especifique um monitor")
	}

	var horarios []model.Horario
	if err := repository.DownloadHorarios(monitor, &horarios); err != nil {
		return err
	}

	c.JSON(http.StatusOK, horarios)
	return nil
}
