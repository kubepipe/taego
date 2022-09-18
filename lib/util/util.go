package util

import (
	"os"

	"taego/mconst"
)

func GetMode() mconst.MODE {
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = string(mconst.MODE_DEBUG)
	}
	return mconst.MODE(mode)
}
