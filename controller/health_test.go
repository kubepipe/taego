package controller

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestOk(t *testing.T) {
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
		Ok(tt.args.c)
	}
}
