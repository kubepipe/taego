package controller

import (
	"context"
	"reflect"
	"taego/lib/merrors"
	"taego/lib/mtrace"
	"taego/mconst"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_success(t *testing.T) {
	type args struct {
		c   *gin.Context
		obj any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		success(tt.args.c, tt.args.obj)
	}
}

func Test_fail(t *testing.T) {
	type args struct {
		c    *gin.Context
		merr merrors.Merr
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		fail(tt.args.c, tt.args.merr)
	}
}

func Test_res(t *testing.T) {
	type args struct {
		c        *gin.Context
		httpcode int
		response *mconst.Response
		data     any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		res(tt.args.c, tt.args.httpcode, tt.args.response, tt.args.data)
	}
}

func Test_traceInfo(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
		want *mconst.TraceInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := traceInfo(tt.args.c); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. traceInfo() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSetSpan(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		SetSpan(tt.args.c)
	}
}

func TestgetSpan(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
		want context.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := getSpan(tt.args.c); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. getSpan() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestgetTrace(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
		want *mtrace.Trace
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := getTrace(tt.args.c); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. getTrace() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_failNot200(t *testing.T) {
	type args struct {
		c        *gin.Context
		httpcode int
		merr     merrors.Merr
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		failNot200(tt.args.c, tt.args.httpcode, tt.args.merr)
	}
}
