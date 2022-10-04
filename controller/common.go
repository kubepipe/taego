package controller

import (
	"context"
	"fmt"

	"taego/lib/merrors"
	"taego/lib/mlog"
	"taego/lib/mtrace"
	"taego/lib/util"
	"taego/mconst"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Ok(c *gin.Context) {
	success(c, "ok")
}

func success(c *gin.Context, obj any) {
	res(c, 200, &mconst.Response{
		Trace: traceInfo(c),
	}, obj)
}

func fail(c *gin.Context, merr merrors.Merr) {
	res(c, 200, &mconst.Response{
		ErrCode: merr.Code(),
		Message: merr.Error(),
		Trace:   traceInfo(c),
	}, nil)
}

func failNot200(c *gin.Context, httpcode int, merr merrors.Merr) {
	res(c, httpcode, &mconst.Response{
		ErrCode: merr.Code(),
		Message: merr.Error(),
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
	if response != nil && response.ErrCode != 0 {
		mlog.Info("fail response",
			zap.Int("httpcode", httpcode),
			zap.String("message", response.Message),
			zap.Int("errcode", response.ErrCode),
			zap.Int32("trace", response.Trace.Id),
		)
	}
	c.Abort()
}

func traceInfo(c *gin.Context) *mconst.TraceInfo {

	span := getSpan(c)

	return &mconst.TraceInfo{
		Id:       mtrace.GetTraceId(span),
		SourceIp: c.ClientIP(),
		ServerIp: util.GetLocalIp(),
	}
}

const spanKey = "mspan"

func SetSpan(c *gin.Context) {
	span, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	// trace
	name := fmt.Sprintf("%s-%s%s",
		c.Request.Method, c.Request.Host, c.Request.RequestURI)
	trace := mtrace.New(name)
	defer trace.Done()

	span = mtrace.ContextWithTrace(span, trace)
	c.Set(spanKey, span)

	c.Next()
}

func getSpan(c *gin.Context) context.Context {
	if ctx, ok := c.Get(spanKey); ok {
		return ctx.(context.Context)
	}
	return context.Background()
}

func getTrace(c *gin.Context) mtrace.Trace {
	return mtrace.GetTrace(getSpan(c))
}
