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

	//TODO 加锁
	mlog.Info(logPrefix + "Try to elect")
	//hostname, err := os.Hostname()
	//if err != nil {
	//	return err
	//}

	const watcherKey = "wukong_watcher"
	//err = mylock.Lock(watcherKey, time.Minute*10, hostname)
	//if err != nil {
	//	return err
	//}

	//TODO 解锁
	defer func() {
		//if err := mylock.Unlock(watcherKey, hostname); err != nil {
		//	mlog.Error(logPrefix, "err", err)
		//}
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
