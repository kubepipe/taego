package api

import (
	ctl "taego/controller"

	"github.com/gin-gonic/gin"
)

func setRoute(e *gin.Engine) {
	e.Any("/", ctl.Ok)
}
