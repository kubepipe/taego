package merrors

import (
	"errors"
	"reflect"
	"testing"
)

func TestMyerr_Error(t *testing.T) {
	tests := []struct {
		name string
		e    Myerr
		want string
	}{
		{"case0", New("case0"), "case0"},
		{"case1", New("case1"), "case1"},
		{"case2", New("case2"), "case2"},
	}
	for _, tt := range tests {
		if got := tt.e.Error(); got != tt.want {
			t.Errorf("%q. Myerr.Error() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNew(t *testing.T) {
	type args struct {
		message string
	}

	tests := []struct {
		name string
		args args
		want Myerr
	}{
		{"case0", args{"case0 msg"}, Myerr("case0 msg")},
		{"case1", args{"case1 msg"}, Myerr("case1 msg")},
		{"case2", args{"case2 msg"}, Myerr("case2 msg")},
	}
	for _, tt := range tests {
		if got := New(tt.args.message); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. New() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNewByError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want Myerr
	}{
		{"case0", args{errors.New("case0 msg")}, Myerr("case0 msg")},
		{"case1", args{errors.New("case1 msg")}, Myerr("case1 msg")},
		{"case2", args{errors.New("case2 msg")}, Myerr("case2 msg")},
	}
	for _, tt := range tests {
		if got := NewByError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewByError() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
