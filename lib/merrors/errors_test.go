package merrors

import (
	"reflect"
	"testing"
)

func TestMyerr_Error(t *testing.T) {
	tests := []struct {
		name string
		e    Myerr
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.e.Error(); got != tt.want {
			t.Errorf("%q. Myerr.Error() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestMyerr_Code(t *testing.T) {
	tests := []struct {
		name string
		e    Myerr
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.e.Code(); got != tt.want {
			t.Errorf("%q. Myerr.Code() = %v, want %v", tt.name, got, tt.want)
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := NewByError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewByError() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
