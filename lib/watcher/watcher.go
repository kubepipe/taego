package watcher

import (
	"sync"
	"time"

	"taego/lib/mlog"

	"go.uber.org/zap"
)

const (
	logPrefix = "mwatcher"
)

func Watch(funcs []func() error) {

	for range time.Tick(time.Second * 5) {
		if err := watch(funcs); err != nil {
			mlog.Error(logPrefix, zap.Error(err))
		}
	}
}

func watch(funcs []func() error) error {
	if len(funcs) == 0 {
		return nil
	}

	//TODO lock

	//TODO unlock
	defer func() {

	}()

	w := &wait{}
	for _, f := range funcs {
		w.run(f)
	}

	w.Wait()
	return nil
}

type wait struct {
	sync.WaitGroup
}

func (wg *wait) run(runner func() error) {
	wg.Add(1)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				mlog.Error(logPrefix, zap.Any("recover", e))
			}
			wg.Done()
		}()

		if err := runner(); err != nil {
			mlog.Error(logPrefix, zap.Error(err))
		}
	}()
}
