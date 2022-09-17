package controller

import (
	"taego/lib/mlog"
	"taego/lib/mtrace"
	"taego/lib/util"
	"taego/mconst"

	"github.com/gin-gonic/gin"
)

func success(c *gin.Context, obj any) {
	res(c, 200, &mconst.Response{
		Success: true,
		Trace:   traceInfo(c),
	}, obj)
}

func fail(c *gin.Context, err error) {
	res(c, 200, &mconst.Response{
		Success: false,
		Message: err.Error(),
		Trace:   traceInfo(c),
	}, nil)
}

func res(c *gin.Context, httpcode int, response *mconst.Response, data any) {
	c.JSON(httpcode, struct {
		*mconst.Response
		Data any `json:"data,omitempty"`
	}{
		Response: response,
		Data:     data,
	})
	if response != nil && !response.Success {
		mlog.Info("fail response ", "httpcode ", httpcode, response.Message, response.Trace.Id)
	}
	c.Abort()
}

func traceInfo(c *gin.Context) *mconst.TraceInfo {

	span := util.GetSpanFromGin(c)

	return &mconst.TraceInfo{
		Id:       mtrace.GetTraceId(span),
		SourceIp: c.ClientIP(),
		ServerIp: util.GetLocalIp(),
	}
}
