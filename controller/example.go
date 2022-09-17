package controller

import (
	"taego/lib/util"
	"taego/service/example"

	"github.com/gin-gonic/gin"
)

func Example(c *gin.Context) {
	// TODO some curd

	// get date from the service
	req := &example.ReqExample{}
	res, err := example.GetExampleData(util.GetSpanFromGin(c), req)
	if err != nil {
		fail(c, err)
		return
	}

	// response
	success(c, string(res))
}
