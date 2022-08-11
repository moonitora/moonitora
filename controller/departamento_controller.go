package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/repository"
	"net/http"
)

func GetDepartamentos(c *gin.Context) error {
	var depts []model.Departamento
	if err := repository.DownloadDepartamentos(&depts); err != nil {
		return err
	}
	c.JSON(http.StatusOK, depts)
	return nil
}
