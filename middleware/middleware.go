package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/authorization"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/repository"
	"github.com/victorbetoni/moonitora/util"
	"net/http"
)

func CheckAuthenticated(forward util.HandleFuncError) util.HandleFuncError {
	return func(c *gin.Context) (int, error) {
		if _, err := authorization.ExtractUser(c); err != nil {
			return http.StatusUnauthorized, errors.New("invalid token")
		}

		forward(c)
		return 0, nil
	}
}

func CheckAdministrator(forward util.HandleFuncError) util.HandleFuncError {
	return func(c *gin.Context) (int, error) {
		user := model.Monitor{}
		email, err := authorization.ExtractUser(c)
		if err != nil {
			return http.StatusUnauthorized, errors.New("invalid token")
		}

		if err := repository.DownloadMonitor(email, &user); err != nil {
			return http.StatusInternalServerError, err
		}

		if user.Adm != 1 {
			return http.StatusUnauthorized, errors.New("unauthorized")
		}
		forward(c)
		return 0, nil
	}
}

func AbortOnError(handler util.HandleFuncError) util.HandleFuncError {
	return func(c *gin.Context) (int, error) {
		if status, err := handler(c); err != nil {
			c.AbortWithStatusJSON(status, gin.H{"status": false, "message": err.Error(), "body": ""})
		}
		return 0, nil
	}
}
