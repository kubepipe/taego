package controller

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestExample(t *testing.T) {
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
		Example(tt.args.c)
	}
}

func Test_demohandle(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		demohandle(tt.args.ctx)
	}
}
