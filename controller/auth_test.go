package controller

import (
	"context"
	"reflect"
	"taego/mconst"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuth(t *testing.T) {
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
		Auth(tt.args.c)
	}
}

func Test_setUser(t *testing.T) {
	type args struct {
		ctx      context.Context
		userinfo *mconst.Userinfo
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		setUser(tt.args.ctx, tt.args.userinfo)
	}
}

func Test_getUser(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *mconst.Userinfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := getUser(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. getUser() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
