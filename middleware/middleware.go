package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/authorization"
	"github.com/victorbetoni/moonitora/util"
	"net/http"
)

func Authorize(forward util.HandleFuncError) util.HandleFuncError {
	return func(c *gin.Context) error {
		if _, err := authorization.ExtractUser(c); err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "invalid token"})
		}

		forward(c)
		return nil
	}
}

func AbortOnError(handler util.HandleFuncError) util.HandleFuncError {
	return func(c *gin.Context) error {
		if err := handler(c); err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		}
		return nil
	}
}
