package controller

import (
	"taego/lib/config"
	"taego/mconst"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("system-token")
	if token != config.Config.UString("system.token", "") {
		fail(c, mconst.ERROR_UNAUTHORIZED)
		return
	}
}
