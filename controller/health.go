package controller

import (
	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context) {
	success(c, "ok")
}

func Health(c *gin.Context) {
	// TODO some check here

	// check db

	success(c, "ok")
}
