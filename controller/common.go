package controller

import (
	"context"
	"fmt"

	"taego/lib/mlog"
	"taego/lib/mtrace"
	"taego/lib/util"
	"taego/mconst"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		mlog.Info("fail response",
			zap.Int("httpcode", httpcode),
			zap.String("message", response.Message),
			zap.Int32("trace", response.Trace.Id),
		)
	}
	c.Abort()
}

func traceInfo(c *gin.Context) *mconst.TraceInfo {

	span := GetSpan(c)

	return &mconst.TraceInfo{
		Id:       mtrace.GetTraceId(span),
		SourceIp: c.ClientIP(),
		ServerIp: util.GetLocalIp(),
	}
}

const spanKey = "mspan"

func GetSpanKey() string {
	return spanKey
}

func SetSpan(c *gin.Context) {
	span := context.Background()

	// trace
	name := fmt.Sprintf("%s-%s%s",
		c.Request.Method, c.Request.Host, c.Request.RequestURI)
	trace := mtrace.New(name)
	defer func() { trace.Done() }()

	span = mtrace.ContextWithTrace(span, trace)
	c.Set(spanKey, span)

	c.Next()
}

func GetSpan(c *gin.Context) context.Context {
	if ctx, ok := c.Get(spanKey); ok {
		return ctx.(context.Context)
	}
	return context.Background()
}

func GetTrace(c *gin.Context) *mtrace.Trace {
	return mtrace.GetTrace(GetSpan(c))
}
