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
	AdminAction    bool
}

var routes = []Route{
	{
		Handler:     controller.Login,
		URI:         "/login",
		RequireAuth: false,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.POST(uri, handler)
		},
		AdminAction: false,
	},
	{
		Handler:     controller.Register,
		URI:         "/register",
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.POST(uri, handler)
		},
		AdminAction: true,
	},
	{
		Handler:     controller.FetchMonitores,
		URI:         "/monitores",
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.GET(uri, handler)
		},
		AdminAction: false,
	},
	{
		Handler:     controller.GetDepartamentos,
		URI:         "/departamentos",
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.GET(uri, handler)
		},
		AdminAction: false,
	},
	{
		Handler:     controller.FetchHorarios,
		URI:         "/horario",
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.GET(uri, handler)
		},
		AdminAction: false,
	},
	{
		Handler:     controller.PostHorario,
		URI:         "/horario",
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.POST(uri, handler)
		},
		AdminAction: false,
	},
	{
		Handler:     controller.PostMonitoria,
		URI:         "/monitoria",
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.POST(uri, handler)
		},
		AdminAction: false,
	},
	{
		Handler:     controller.DeleteHorario,
		URI:         "horario",
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.DELETE(uri, handler)
		},
		AdminAction: false,
	},
	{
		Handler:     controller.FetchMonitorias,
		URI:         "/monitorias",
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.GET(uri, handler)
		},
		AdminAction: false,
	},
	{
		Handler:     controller.CheckDisponibility,
		URI:         "/checkdisponibility",
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.POST(uri, handler)
		},
		AdminAction: false,
	},
	{
		Handler:     controller.CancelMonitoria,
		URI:         "/monitoria/cancel",
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.GET(uri, handler)
		},
		AdminAction: false,
	},
	{
		Handler:     controller.ConfirmMonitoria,
		URI:         "/monitoria/confirm",
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.GET(uri, handler)
		},
		AdminAction: false,
	},
	{
		Handler:     controller.ConcludeMonitoria,
		URI:         "/monitoria/conclude",
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.GET(uri, handler)
		},
		AdminAction: false,
	},
	{
		Handler:     controller.PostDepartamento,
		URI:         "/departamento",
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.POST(uri, handler)
		},
		AdminAction: true,
	},
	{
		Handler:     controller.DeleteDepartamento,
		URI:         "/departamento",
		RequireAuth: true,
		AssignFunction: func(e *gin.Engine, handler gin.HandlerFunc, uri string) {
			e.DELETE(uri, handler)
		},
		AdminAction: true,
	},
}

func Setup(e *gin.Engine) *gin.Engine {
	var r []Route
	r = append(r, routes...)
	for _, x := range r {
		if x.RequireAuth {
			if x.AdminAction {
				Assign(x.AssignFunction, middleware.AbortOnError(middleware.CheckAuthenticated(middleware.CheckAdministrator(x.Handler))), x.URI, e)
			} else {
				Assign(x.AssignFunction, middleware.AbortOnError(middleware.CheckAuthenticated(x.Handler)), x.URI, e)
			}
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
