package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/authorization"
	"github.com/victorbetoni/moonitora/database"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/security"
	"net/http"
)

func Login(c *gin.Context) error {
	loginInfo := model.Login{}
	if err := c.BindJSON(&loginInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request", "user": "", "jwt": ""})
		return err
	}

	db := database.GrabDB()
	var userLogin model.Login
	if err := db.Get(&userLogin, "SELECT * FROM login WHERE email=$1", loginInfo.Email); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "Usuário não encontrado", "user": "", "jwt": ""})
			return err
		}
	}

	if err := security.ComparePassword(userLogin.Password, loginInfo.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "Senha incorreta", "user": "", "jwt": ""})
		return err
	}

	token := authorization.GenerateToken(userLogin.Email)
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Login efetuado com sucesso", "user": userLogin.Email, "jwt": token})
	return nil
}
