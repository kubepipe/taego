package api

import (
	"context"
	"fmt"
	"os"
	"time"

	"taego/lib/util"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	gin.SetMode(os.Getenv("MODE"))
	e := gin.New()
	e.Use(setspan)
	e.Use(gin.LoggerWithFormatter(ginLoger))
	e.Use(gin.Recovery())

	setRoute(e)
	return e
}

func setspan(c *gin.Context) {
	// TODO
	//trace.SetSpan(c, config.OpentraceSwitch())
}

func ginLoger(param gin.LogFormatterParams) string {

	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}

	var traceId string
	if span, ok := param.Keys["span"]; ok {
		spanc, _ := span.(context.Context)
		traceId = util.GetTraceId(spanc)
	}

	return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v | %s\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		traceId,
		param.ErrorMessage,
	)
}
