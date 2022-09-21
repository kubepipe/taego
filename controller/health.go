package controller

import (
	"net/http"
	"time"

	"taego/lib/merrors"
	"taego/lib/mmysql"
	"taego/lib/util"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context) {
	success(c, "ok")
}

func Health(c *gin.Context) {
	// TODO some check here

	// check db
	checkDB := func() error {
		return mmysql.Ping()
	}
	if err := util.RetryUntilSuccess(checkDB, 3, time.Millisecond*100); err != nil {
		failNot200(c, http.StatusInternalServerError, merrors.New(err))
		return
	}

	success(c, "ok")
}
