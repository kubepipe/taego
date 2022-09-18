package mlog

import (
	"testing"

	"go.uber.org/zap"
)

func TestDebug(t *testing.T) {
	type args struct {
		message string
		args    []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Debug(tt.args.message, tt.args.args...)
	}
}

func TestDebugf(t *testing.T) {
	type args struct {
		template string
		args     []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Debugf(tt.args.template, tt.args.args...)
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		message string
		args    []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Info(tt.args.message, tt.args.args...)
	}
}

func TestInfof(t *testing.T) {
	type args struct {
		template string
		args     []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Infof(tt.args.template, tt.args.args...)
	}
}

func TestWarn(t *testing.T) {
	type args struct {
		message string
		args    []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Warn(tt.args.message, tt.args.args...)
	}
}

func TestWarnf(t *testing.T) {
	type args struct {
		template string
		args     []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Warnf(tt.args.template, tt.args.args...)
	}
}

func TestError(t *testing.T) {
	type args struct {
		message string
		args    []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Error(tt.args.message, tt.args.args...)
	}
}

func TestErrorf(t *testing.T) {
	type args struct {
		template string
		args     []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Errorf(tt.args.template, tt.args.args...)
	}
}

func TestFatal(t *testing.T) {
	type args struct {
		message string
		args    []zap.Field
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Fatal(tt.args.message, tt.args.args...)
	}
}

func TestFatalf(t *testing.T) {
	type args struct {
		template string
		args     []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Fatalf(tt.args.template, tt.args.args...)
	}
}
