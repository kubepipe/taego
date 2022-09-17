package api

import (
	"os"

	ctl "taego/controller"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	gin.SetMode(os.Getenv("MODE"))
	e := gin.New()
	e.Use(ctl.SetSpan)
	e.Use(gin.Recovery())

	setRoute(e)
	return e
}
