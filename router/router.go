package router

import "github.com/gin-gonic/gin"

type Route struct {
	RequireAuth bool
	Handler     func(c *gin.Context) error
	URI         string
}

func Build() *gin.Engine {
	return nil
}
