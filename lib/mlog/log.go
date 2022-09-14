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

func Debug(args ...interface{}) {
	l.Debug(args)
}

func Info(args ...interface{}) {
	l.Info(args)
}

func Warn(args ...interface{}) {
	l.Warn(args)
}

func Error(args ...interface{}) {
	l.Error(args)
}

func Fatal(args ...interface{}) {
	l.Fatal(args)
}
