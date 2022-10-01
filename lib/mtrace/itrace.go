package mtrace

import "go.uber.org/zap"

type Trace interface {
	// return traceid
	GetTraceId() int32

	// You can count the total time spent here, print logs, and upload monitoring metrics
	Done(args ...zap.Field)

	// generate child trace
	SubTrace(string) Trace

	// log with traceid
	Log(string, ...zap.Field)
}
