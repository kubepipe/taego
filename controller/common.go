package controller

import (
	"context"

	"taego/lib/config"
	"taego/lib/mlog"
	"taego/lib/trace"
	"taego/lib/util"
	"taego/mconst"

	"github.com/gin-gonic/gin"
)

func success(c *gin.Context, obj interface{}) {
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

func unauth(c *gin.Context) {
	res(c, 200, &mconst.Response{
		Message: "unauthorized",
		Success: false,
		Trace:   traceInfo(c),
	}, nil)
}

func res(c *gin.Context, httpcode int, response *mconst.Response, data interface{}) {
	c.JSON(httpcode, struct {
		*mconst.Response
		Data interface{} `json:"data,omitempty"`
	}{
		Response: response,
		Data:     data,
	})
	if response != nil && !response.Success {
		var args = make([]interface{}, 0, 5)
		args = []interface{}{"fail response", "httpcode", httpcode, response.Message}
		if response.Trace != nil {
			args = append(args, response.Trace.Id)
		}
		mlog.Info(args)
	}
	c.Abort()
}

// TODO trace
func traceInfo(c *gin.Context) *trace.Trace {
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

	return &trace.Trace{
		Id:       traceId.(string),
		SrcIp:    c.ClientIP(),
		ServerIp: util.GetLocalIp(),
	}
}

// TODO
func GetSpan(c *gin.Context) context.Context {
	return context.TODO()
}
