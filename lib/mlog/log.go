package mlog

import (
	"log"
	"os"

	"go.uber.org/zap"
)

var l *zap.SugaredLogger

func init() {
	var (
		logp *zap.Logger
		err  error
	)

	options := []zap.Option{
		zap.AddCallerSkip(1),
	}
	switch os.Getenv("MODE") {
	case "release":
		logp, err = zap.NewProduction(options...)
	default:
		logp, err = zap.NewDevelopment(options...)
	}

	if err != nil {
		log.Fatal(err)
	}
	l = logp.Sugar()
}

func Debug(args ...any) {
	l.Debug(args)
}

func Debugf(template string, args ...any) {
	l.Debugf(template, args...)
}

func Info(args ...any) {
	l.Info(args...)
}

func Infof(template string, args ...any) {
	l.Infof(template, args...)
}

func Warn(args ...any) {
	l.Warn(args...)
}

func Warnf(template string, args ...any) {
	l.Warnf(template, args...)
}

func Error(args ...any) {
	l.Error(args...)
}

func Errorf(template string, args ...any) {
	l.Errorf(template, args...)
}

func Fatal(args ...any) {
	l.Fatal(args...)
}

func Fatalf(template string, args ...any) {
	l.Fatalf(template, args...)
}
