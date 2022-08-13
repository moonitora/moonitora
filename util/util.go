package util

import "github.com/gin-gonic/gin"

type HandleFuncError func(c *gin.Context) (int, error)
type AssignFunction func(e *gin.Engine, handler gin.HandlerFunc, uri string)
