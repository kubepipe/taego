package util

import (
	"reflect"
	"taego/mconst"
	"testing"
)

func TestGetMode(t *testing.T) {
	tests := []struct {
		name string
		want mconst.MODE
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := GetMode(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. GetMode() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
