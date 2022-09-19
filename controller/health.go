package controller

import (
	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context) {
	success(c, "ok")
}

func Health(c *gin.Context) {
	// TODO some check here
	//if !ok {
	//	failNot200(c, http.StatusInternalServerError, merrors.Get(merrors.ERROR_UNHEALTHY))
	//}

	success(c, "ok")
}
