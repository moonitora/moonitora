package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/authorization"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/repository"
	"net/http"
	"strconv"
)

type IncomingUser struct {
	Monitor model.Monitor `json:"monitor"`
	Login   model.Login   `json:"login"`
}

type Response struct {
	JWT     string        `json:"jwt"`
	Monitor model.Monitor `json:"monitor"`
}

func FetchMonitores(c *gin.Context) error {
	dept, ok := c.GetQuery("departamento")
	if !ok {
		return nil
	}

	val, err := strconv.Atoi(dept)
	if err != nil {
		return err
	}

	var monitores []model.Monitor
	if err := repository.DownloadMonitores(val, &monitores); err != nil {
		return err
	}

	c.JSON(http.StatusOK, monitores)
	return nil
}

func Register(c *gin.Context) error {
	incoming := IncomingUser{}
	if err := c.BindJSON(&incoming); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request", "user": "", "jwt": ""})
		return err
	}

	if err := repository.InsertMonitor(incoming.Monitor); err != nil {
		return err
	}

	if err := repository.InsertLogin(incoming.Login); err != nil {
		return err
	}

	token := authorization.GenerateToken(incoming.Monitor.Email)
	c.JSON(http.StatusOK, Response{JWT: token, Monitor: incoming.Monitor})
	return nil
}
