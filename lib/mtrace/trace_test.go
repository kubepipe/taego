package mtrace

import (
	"reflect"
	"testing"

	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want Trace
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := New(tt.args.name); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. New() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_trace_SubTrace(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		tr   *trace
		args args
		want Trace
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.tr.SubTrace(tt.args.name); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. trace.SubTrace() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_trace_Log(t *testing.T) {
	type args struct {
		message string
		args    []zap.Field
	}
	tests := []struct {
		name string
		tr   *trace
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.tr.Log(tt.args.message, tt.args.args...)
	}
}

func Test_trace_Done(t *testing.T) {
	tests := []struct {
		name string
		tr   *trace
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.tr.Done()
	}
}

func Test_trace_GetTraceId(t *testing.T) {
	tests := []struct {
		name string
		tr   *trace
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.tr.GetTraceId(); got != tt.want {
			t.Errorf("%q. trace.GetTraceId() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
