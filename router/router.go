package router

import (
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/controller"
	"github.com/victorbetoni/moonitora/middleware"
)

type HandleFuncError func(c *gin.Context) error
type AssignFunction func(e *gin.Engine, handler gin.HandlerFunc, uri string)

type Route struct {
	RequireAuth    bool
	Handler        HandleFuncError
	URI            string
	AssignFunction AssignFunction
}

var loginRoutes = []Route{
	{
		Handler:     controller.Login,
		URI:         "/login",
		RequireAuth: false,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.POST(uri, handler)
		},
	},
	{
		Handler:     controller.Register,
		URI:         "/register",
		RequireAuth: false,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.POST(uri, handler)
		},
	},
}

func Setup(e *gin.Engine) *gin.Engine {
	var routes []Route
	routes = append(routes, loginRoutes...)
	for _, x := range routes {
		if x.RequireAuth {
			Assign(x.AssignFunction, middleware.AbortOnError(middleware.Authorize(x.Handler)), x.URI, e)
		} else {
			Assign(x.AssignFunction, middleware.AbortOnError(x.Handler), x.URI, e)
		}
		Assign(x.AssignFunction, x.Handler, x.URI, e)
	}

	return nil
}

func Assign(assign AssignFunction, handler HandleFuncError, uri string, router *gin.Engine) {
	assign(router, func(context *gin.Context) {
		handler(context)
	}, uri)
}
