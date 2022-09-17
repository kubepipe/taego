package mlog

import (
	"log"
	"os"

	"go.uber.org/zap"
)

var l *zap.Logger

func init() {
	var err error

	options := []zap.Option{
		zap.AddCallerSkip(1),
	}
	switch os.Getenv("MODE") {
	case "release":
		l, err = zap.NewProduction(options...)
	default:
		l, err = zap.NewDevelopment(options...)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func Debug(message string, args ...zap.Field) {
	l.Debug(message, args...)
}

func Debugf(template string, args ...any) {
	l.Sugar().Debugf(template, args...)
}

func Info(message string, args ...zap.Field) {
	l.Info(message, args...)
}

func Infof(template string, args ...any) {
	l.Sugar().Infof(template, args...)
}

func Warn(message string, args ...zap.Field) {
	l.Warn(message, args...)
}

func Warnf(template string, args ...any) {
	l.Sugar().Warnf(template, args...)
}

func Error(message string, args ...zap.Field) {
	l.Error(message, args...)
}

func Errorf(template string, args ...any) {
	l.Sugar().Errorf(template, args...)
}

func Fatal(message string, args ...zap.Field) {
	l.Fatal(message, args...)
}

func Fatalf(template string, args ...any) {
	l.Sugar().Fatalf(template, args...)
}
