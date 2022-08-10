package router

import (
	"github.com/gin-gonic/gin"
	"github.com/victorbetoni/moonitora/controller"
	"github.com/victorbetoni/moonitora/middleware"
	"github.com/victorbetoni/moonitora/util"
	"net/http"
)

type Route struct {
	RequireAuth    bool
	Handler        util.HandleFuncError
	URI            string
	AssignFunction util.AssignFunction
}

var routes = []Route{
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
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.POST(uri, handler)
		},
	},
	{
		Handler:     controller.FetchMonitores,
		URI:         "/monitores",
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.GET(uri, handler)
		},
	},
}

func Setup(e *gin.Engine) *gin.Engine {
	var routes []Route
	routes = append(routes, routes...)
	for _, x := range routes {
		if x.RequireAuth {
			Assign(x.AssignFunction, middleware.AbortOnError(middleware.Authorize(x.Handler)), x.URI, e)
		} else {
			Assign(x.AssignFunction, middleware.AbortOnError(x.Handler), x.URI, e)
		}
	}
	e.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Pong")
	})
	return e
}

func Assign(assign util.AssignFunction, handler util.HandleFuncError, uri string, router *gin.Engine) {
	assign(router, func(context *gin.Context) {
		handler(context)
	}, uri)
}
