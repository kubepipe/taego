package api

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_setRoute(t *testing.T) {
	type args struct {
		e *gin.Engine
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		setRoute(tt.args.e)
	}
}
