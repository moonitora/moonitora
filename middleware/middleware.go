package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/authorization"
	"github.com/victorbetoni/moonitora/model"
	"github.com/victorbetoni/moonitora/repository"
	"github.com/victorbetoni/moonitora/util"
	"net/http"
)

func CheckAuthenticated(forward util.HandleFuncError) util.HandleFuncError {
	return func(c *gin.Context) error {
		if _, err := authorization.ExtractUser(c); err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "invalid token"})
		}

		forward(c)
		return nil
	}
}

func CheckAdministrator(forward util.HandleFuncError) util.HandleFuncError {
	return func(c *gin.Context) error {
		user := model.Monitor{}
		var email string
		if a, err := authorization.ExtractUser(c); err != nil {
			email = a
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "invalid token"})
		}

		if err := repository.DownloadMonitor(email, &user); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": err.Error()})
		}

		if user.Adm != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "unauthorized"})
		}

		forward(c)
		return nil
	}
}

func AbortOnError(handler util.HandleFuncError) util.HandleFuncError {
	return func(c *gin.Context) error {
		if err := handler(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
