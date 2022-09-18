package util

import "testing"

func TestGetFuncName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := GetFuncName(); got != tt.want {
			t.Errorf("%q. GetFuncName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
