package mtrace

import "context"

type traceKey struct{}

func GetTrace(ctx context.Context) Trace {
	if v, ok := ctx.Value(traceKey{}).(Trace); ok {
		return v
	}
	// avoid panic
	return &trace{}
}

func SubTrace(ctx context.Context, name string) Trace {
	return GetTrace(ctx).SubTrace(name)
}

func GetTraceId(ctx context.Context) int32 {
	return GetTrace(ctx).GetTraceId()
}

func ContextWithTrace(ctx context.Context, t Trace) context.Context {
	return context.WithValue(ctx, traceKey{}, t)
}
