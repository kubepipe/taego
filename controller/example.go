package controller

import (
	"context"
	"taego/lib/merrors"
	"taego/lib/mtrace"
	"taego/service/example"
	"time"

	"github.com/gin-gonic/gin"
)

func Example(c *gin.Context) {

	demohandle(GetSpan(c))

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

func demohandle(ctx context.Context) {
	trace := mtrace.SubTrace(ctx, "demohandle")
	defer func() { trace.Done() }()

	trace.Log("i am doing some curd")
	time.Sleep(time.Microsecond * 10)
}
