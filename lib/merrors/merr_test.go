package merrors

import (
	"errors"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want Merr
	}{
		// TODO: Add test cases.
		{"case0", args{errors.New("case0")}, merr("case0")},
	}
	for _, tt := range tests {
		if got := New(tt.args.err); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. New() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGet(t *testing.T) {
	type args struct {
		errcode Code
	}
	tests := []struct {
		name string
		args args
		want Merr
	}{
		// TODO: Add test cases.
		{"unhealthy", args{ERROR_UNHEALTHY}, merr(codemap[ERROR_UNHEALTHY])},
		{"unauthorized", args{ERROR_UNAUTHORIZED}, merr(codemap[ERROR_UNAUTHORIZED])},
	}
	for _, tt := range tests {
		if got := Get(tt.args.errcode); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Get() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_merr_Error(t *testing.T) {
	tests := []struct {
		name string
		e    merr
		want string
	}{
		// TODO: Add test cases.
		{"case0", merr("case0"), "case0"},
		{"unhealth", merr(codemap[ERROR_UNHEALTHY]), codemap[ERROR_UNHEALTHY]},
		{"unauth", merr(codemap[ERROR_UNAUTHORIZED]), codemap[ERROR_UNAUTHORIZED]},
		{"unknow", merr("taego test merr"), "taego test merr"},
	}
	for _, tt := range tests {
		if got := tt.e.Error(); got != tt.want {
			t.Errorf("%q. merr.Error() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_merr_Code(t *testing.T) {
	tests := []struct {
		name string
		e    merr
		want int
	}{
		// TODO: Add test cases.
		{"unhealth", merr(codemap[ERROR_UNHEALTHY]), int(ERROR_UNHEALTHY)},
		{"unauth", merr(codemap[ERROR_UNAUTHORIZED]), int(ERROR_UNAUTHORIZED)},
		{"unknow", merr("taego test merr"), -1},
	}
	for _, tt := range tests {
		if got := tt.e.Code(); got != tt.want {
			t.Errorf("%q. merr.Code() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
