package util

import (
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

func QueryInt(c *gin.Context, key string) int {
	if c == nil {
		return 0
	}
	v, _ := strconv.Atoi(c.Query(key))
	return v
}
