package controller

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context) {
	success(c, time.Now().String())
}
