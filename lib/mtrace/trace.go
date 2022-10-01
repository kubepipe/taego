package mtrace

import (
	"math/rand"
	"sync"
	"time"

	"taego/lib/mlog"

	"go.uber.org/zap"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type trace struct {
	id        int32
	name      string
	startTime time.Time

	sync.Once
}

func New(name string) Trace {
	return &trace{
		id:        rand.Int31n(int32(1<<31-1)) + 1,
		name:      name,
		startTime: time.Now(),
	}
}

// use the same id as parent trace
func (t *trace) SubTrace(name string) Trace {
	return &trace{
		id:        t.id,
		name:      name,
		startTime: time.Now(),
	}
}

// mlog.Info with trace.id
func (t *trace) Log(message string, args ...zap.Field) {
	traces := []zap.Field{
		zap.Int32("trace", t.id),
		zap.String("traceName", t.name),
	}
	args = append(args, traces...)
	mlog.Info(message, args...)
}

// trace finish, caculate total time
func (t *trace) Done(args ...zap.Field) {
	t.Do(func() {
		args = append(args,
			zap.Int32("trace", t.id),
			zap.String("traceName", t.name),
			zap.String("totalTime", time.Now().Sub(t.startTime).String()),
		)
		mlog.Info("step done", args...)
	})
}

func (t *trace) GetTraceId() int32 {
	return t.id
}
