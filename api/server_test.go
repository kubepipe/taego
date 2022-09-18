package api

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name string
		want *gin.Engine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Run(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Run() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
