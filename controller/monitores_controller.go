package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/authorization"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/repository"
	"net/http"
)

type IncomingUser struct {
	Monitor model.Monitor `json:"monitor"`
	Login   model.Login   `json:"login"`
}

func FetchMonitores(c *gin.Context) (int, error) {
	dept, ok := c.GetQuery("departamento")
	if !ok {
		return http.StatusBadRequest, errors.New("especifique um departamento")
	}

	var monitores []model.Monitor
	if err := repository.DownloadMonitores(dept, &monitores); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "", "body": monitores})
	return 0, nil
}

func Register(c *gin.Context) (int, error) {
	incoming := IncomingUser{}
	if err := c.BindJSON(&incoming); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	if err := repository.InsertMonitor(incoming.Monitor); err != nil {
		return http.StatusBadRequest, err
	}

	if err := repository.InsertLogin(incoming.Login); err != nil {
		return http.StatusBadRequest, err
	}

	token := authorization.GenerateToken(incoming.Monitor.Email)

	c.JSON(http.StatusOK, gin.H{"jwt": token, "body": incoming.Monitor, "status": true, "message": "Usuário cadastrado com sucesso"})
	return 0, nil
}
