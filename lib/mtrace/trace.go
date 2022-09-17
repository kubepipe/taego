package mtrace

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"taego/lib/mlog"

	"go.uber.org/zap"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Trace struct {
	id        int32
	name      string
	startTime time.Time

	sync.Once
}

func New(name string) *Trace {
	return &Trace{
		id:        rand.Int31(),
		name:      name,
		startTime: time.Now(),
	}
}

// use the same id as parent trace
func (t *Trace) subTrace(name string) *Trace {
	return &Trace{
		id:        t.id,
		name:      name,
		startTime: time.Now(),
	}
}

// mlog.Info with trace.id
func (t *Trace) Log(message string, args ...zap.Field) {
	traces := []zap.Field{
		zap.Int32("trace", t.id),
		zap.String("traceName", t.name),
	}
	args = append(args, traces...)
	mlog.Info(message, args...)
}

// trace finish, caculate total time
func (t *Trace) Done() {
	t.Do(func() {
		mlog.Info("step done",
			zap.Int32("trace", t.id),
			zap.String("traceName", t.name),
			zap.String("totalTime", time.Now().Sub(t.startTime).String()),
		)
	})
}

func SubTrace(ctx context.Context, name string) *Trace {
	return GetTrace(ctx).subTrace(name)
}

type traceKey struct{}

func GetTrace(ctx context.Context) *Trace {
	if v, ok := ctx.Value(traceKey{}).(*Trace); ok {
		return v
	}
	// avoid panic
	return &Trace{}
}

func GetTraceId(ctx context.Context) int32 {
	return GetTrace(ctx).id
}

func ContextWithTrace(ctx context.Context, t *Trace) context.Context {
	return context.WithValue(ctx, traceKey{}, t)
}
