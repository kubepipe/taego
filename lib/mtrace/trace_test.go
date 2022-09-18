package mtrace

import (
	"context"
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
		want *Trace
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := New(tt.args.name); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. New() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestTrace_subTrace(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		tr   *Trace
		args args
		want *Trace
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.tr.subTrace(tt.args.name); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Trace.subTrace() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestTrace_Log(t *testing.T) {
	type args struct {
		message string
		args    []zap.Field
	}
	tests := []struct {
		name string
		tr   *Trace
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.tr.Log(tt.args.message, tt.args.args...)
	}
}

func TestTrace_Done(t *testing.T) {
	tests := []struct {
		name string
		tr   *Trace
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.tr.Done()
	}
}

func TestSubTrace(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name string
		args args
		want *Trace
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := SubTrace(tt.args.ctx, tt.args.name); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. SubTrace() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetTrace(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *Trace
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := GetTrace(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. GetTrace() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetTraceId(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := GetTraceId(tt.args.ctx); got != tt.want {
			t.Errorf("%q. GetTraceId() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestContextWithTrace(t *testing.T) {
	type args struct {
		ctx context.Context
		t   *Trace
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := ContextWithTrace(tt.args.ctx, tt.args.t); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ContextWithTrace() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
