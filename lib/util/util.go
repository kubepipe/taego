package util

import (
	"os"
	"time"

	"taego/mconst"
)

func GetMode() mconst.MODE {
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = string(mconst.MODE_DEBUG)
	}
	return mconst.MODE(mode)
}

func RetryUntilSuccess(f func() error, num int, interval time.Duration) error {
	var err error
	for i := 0; i < num; i++ {
		err = f()
		if err == nil {
			return nil
		}
		time.Sleep(interval)
	}
	return err
}
