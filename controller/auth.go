package controller

import (
	"taego/lib/config"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("system-token")
	if token != config.Config.UString("system.token", "") {
		unauth(c)
		return
	}
}
