package util

import (
	"context"
	"os"
	"strconv"

	"taego/mconst"

	"github.com/gin-gonic/gin"
)

func GetMode() mconst.MODE {
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = string(mconst.MODE_DEBUG)
	}
	return mconst.MODE(mode)
}

// TODO
func GetTraceId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	obj := ctx.Value("traceId")
	if obj == nil {
		return ""
	}
	traceId, _ := obj.(string)
	return traceId
}

func QueryInt(c *gin.Context, key string) int {
	if c == nil {
		return 0
	}
	v, _ := strconv.Atoi(c.Query(key))
	return v
}
