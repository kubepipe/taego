package mtrace

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"taego/lib/mlog"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Trace struct {
	id        int32
	name      string
	startTime time.Time
}

func New(name string) *Trace {
	return &Trace{
		id:        rand.Int31(),
		name:      name,
		startTime: time.Now(),
	}
}

// use the same id as parent trace
func (t *Trace) SubTrace(name string) *Trace {
	return &Trace{
		id:        t.id,
		name:      name,
		startTime: time.Now(),
	}
}

// mlog.Info with trace.id
func (t *Trace) Log(args ...any) {
	if t.id != 0 {
		args = append(args, fmt.Sprintf(" trace[%d] [%s]", t.id, t.name))
	}
	mlog.Info(args...)
}

func (t *Trace) Done() {
	// TODO upload metrix
	mlog.Infof("trace[%d] [%s] time[%v]", t.id, t.name, time.Now().Sub(t.startTime))
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
