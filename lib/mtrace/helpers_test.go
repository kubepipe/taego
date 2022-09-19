package mtrace

import (
	"context"
	"reflect"
	"testing"
)

func TestGetTrace(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want Trace
	}{
		// TODO: Add test cases.
		{"case0", args{context.WithValue(context.Background(), traceKey{}, &trace{id: 0})}, &trace{id: 0}},
		{"case1", args{context.WithValue(context.Background(), traceKey{}, &trace{id: 1})}, &trace{id: 1}},
		{"case2", args{context.WithValue(context.Background(), traceKey{}, &trace{id: 2})}, &trace{id: 2}},
	}
	for _, tt := range tests {
		if got := GetTrace(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. GetTrace() = %v, want %v", tt.name, got, tt.want)
		}
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
		want Trace
	}{
		// TODO: Add test cases.
		{"case0", args{context.WithValue(context.Background(), traceKey{}, &trace{id: 100}), "sub0"}, &trace{id: 100, name: "sub0"}},
		{"case1", args{context.WithValue(context.Background(), traceKey{}, &trace{id: 1000}), "sub1"}, &trace{id: 1000, name: "sub1"}},
		{"case2", args{context.WithValue(context.Background(), traceKey{}, &trace{id: 10000}), "sub2"}, &trace{id: 10000, name: "sub2"}},
	}
	for _, tt := range tests {
		if got := SubTrace(tt.args.ctx, tt.args.name); got.GetTraceId() != tt.want.GetTraceId() {
			t.Errorf("%q. SubTrace() = %v, want %v", tt.name, got, tt.want)
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
		{"case0", args{context.WithValue(context.Background(), traceKey{}, &trace{id: 0})}, 0},
		{"case1", args{context.WithValue(context.Background(), traceKey{}, &trace{id: 1})}, 1},
		{"case2", args{context.WithValue(context.Background(), traceKey{}, &trace{id: 2})}, 2},
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
		t   Trace
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		// TODO: Add test cases.
		{"case0", args{
			context.Background(),
			&trace{id: 0},
		}, context.WithValue(context.Background(), traceKey{}, &trace{id: 0})},
		{"case1", args{
			context.Background(),
			&trace{id: 1},
		}, context.WithValue(context.Background(), traceKey{}, &trace{id: 1})},
		{"case2", args{
			context.Background(),
			&trace{id: 2},
		}, context.WithValue(context.Background(), traceKey{}, &trace{id: 2})},
	}
	for _, tt := range tests {
		if got := ContextWithTrace(tt.args.ctx, tt.args.t); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ContextWithTrace() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
