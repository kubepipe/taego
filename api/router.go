package api

import (
	ctl "taego/controller"

	"github.com/gin-gonic/gin"
)

func setRoute(e *gin.Engine) {
	e.Any("/", ctl.Ok)
	e.GET("/health", ctl.Health)

	e.GET("/demo", ctl.Demo)

	v1 := e.Group("/api/v1", ctl.Auth)

	v1.GET("/example", ctl.Example)
}
