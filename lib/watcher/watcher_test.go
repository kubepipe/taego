package watcher

import "testing"

func TestWatch(t *testing.T) {
	type args struct {
		funcs []func() error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Watch(tt.args.funcs)
	}
}

func Test_watch(t *testing.T) {
	type args struct {
		funcs []func() error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := watch(tt.args.funcs); (err != nil) != tt.wantErr {
			t.Errorf("%q. watch() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Test_wait_run(t *testing.T) {
	type args struct {
		runner func() error
	}
	tests := []struct {
		name string
		wg   *wait
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.wg.run(tt.args.runner)
	}
}
