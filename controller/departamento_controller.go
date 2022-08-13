package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/repository"
	"net/http"
)

func GetDepartamentos(c *gin.Context) (int, error) {
	var depts []model.Departamento
	if err := repository.DownloadDepartamentos(&depts); err != nil {
		return http.StatusInternalServerError, err
	}
	c.JSON(http.StatusOK, depts)
	return 0, nil
}
