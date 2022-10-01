package controller

import (
	"taego/dao"
	"taego/lib/merrors"
	"taego/lib/mlog"
	"taego/service/example"

	"github.com/gin-gonic/gin"
)

func Example(c *gin.Context) {

	// db
	us := dao.GetUserNames(GetSpan(c))
	mlog.Infof("get user name: %v", us)

	// get date from the service
	req := &example.ReqExample{}
	res, err := example.GetExampleData(GetSpan(c), req)
	if err != nil {
		fail(c, merrors.New(err))
		return
	}

	// mlog.Info does not contain trace
	GetTrace(c).Log("some other things to do")

	// response
	success(c, string(res))
}
