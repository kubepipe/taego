package util

import (
	"context"
	"fmt"

	"taego/lib/mtrace"

	"github.com/gin-gonic/gin"
)

const spanKey = "mspan"

func SetSpan2Gin(c *gin.Context) {
	trace := mtrace.New(fmt.Sprintf("%s-%s-%s%s",
		c.ClientIP(), c.Request.Method, c.Request.Host, c.Request.RequestURI))

	c.Set(spanKey, mtrace.ContextWithTrace(context.Background(), trace))
}

func GetSpanFromGin(c *gin.Context) context.Context {
	if ctx, ok := c.Get(spanKey); ok {
		return ctx.(context.Context)
	}
	return context.Background()
}
