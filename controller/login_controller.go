package controller

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/authorization"
	"github.com/victorbetoni/moonitora/database"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/security"
	"net/http"
)

func Login(c *gin.Context) (int, error) {
	loginInfo := model.Login{}
	if err := c.BindJSON(&loginInfo); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	db := database.GrabDB()
	var userLogin model.Login
	if err := db.Get(&userLogin, "SELECT * FROM login WHERE email=$1", loginInfo.Email); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("usuario nao encontrado")
		}
	}

	if err := security.ComparePassword(userLogin.Password, loginInfo.Password); err != nil {
		return http.StatusUnauthorized, errors.New("senha incorreta")
	}

	token := authorization.GenerateToken(userLogin.Email)
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Login efetuado com sucesso", "body": userLogin.Email, "jwt": token})
	return 0, nil
}
