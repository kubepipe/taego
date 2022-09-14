package controller

import (
	"context"
	"time"

	"taego/lib/config"
	"taego/lib/mlog"
	"taego/mconst"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context) {
	success(c, time.Now().String())
}

func success(c *gin.Context, obj interface{}) {
	res(c, 200, &mconst.Response{
		Success: true,
		Trace:   traceInfo(c),
	}, obj)
}

func fail(c *gin.Context, err error) {
	res(c, 200, &mconst.Response{
		Code:    -1,
		Success: false,
		Message: err.Error(),
		Trace:   traceInfo(c),
	}, nil)
}

func unauth(c *gin.Context) {
	res(c, 200, &mconst.Response{
		Code:    401,
		Message: "认证未通过",
		Success: false,
		Trace:   traceInfo(c),
	}, nil)
}

func res(c *gin.Context, httpcode int, response *mconst.Response, data interface{}) {
	c.JSON(httpcode, struct {
		mconst.Response
		Data interface{} `json:"data,omitempty"`
	}{
		Response: *response,
		Data:     data,
	})
	if response != nil && !response.Success {
		mlog.Info("fail response", "httpcode", httpcode, response.Message, response.Trace.Id)
	}
	c.Abort()
}

func traceInfo(c *gin.Context) *mconst.Trace {
	if !config.OpentraceSwitch() {
		return nil
	}

	spanC, _ := c.Value("span").(context.Context)
	if spanC == nil {
		return nil
	}
	traceId := spanC.Value("traceId")
	if traceId == nil {
		return nil
	}

	return &mconst.Trace{
		Id:    traceId.(string),
		SrcIp: c.ClientIP(),
	}
}

// TODO
func GetSpan(c *gin.Context) context.Context {
	return context.TODO()
}