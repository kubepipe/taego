package util

import "testing"

func TestHashStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := HashStr(tt.args.str); got != tt.want {
			t.Errorf("%q. HashStr() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
