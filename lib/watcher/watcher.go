package watcher

import (
	"sync"
	"time"

	"taego/lib/mlog"
)

const (
	logPrefix = "miramar-watcher "
)

func Watch(funcs []func() error) {

	for range time.Tick(time.Second * 5) {
		if err := watch(funcs); err != nil {
			mlog.Error(logPrefix, "err", err)
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

	mlog.Info(logPrefix + "Started")

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
				mlog.Error(logPrefix, "err", e)
			}
			wg.Done()
		}()

		if err := runner(); err != nil {
			mlog.Error(logPrefix, "err", err)
		}
	}()
}
