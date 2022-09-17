package main

import (
	"net/http"
	"time"

	"taego/api"
	"taego/lib/config"
	"taego/lib/mlog"

	"github.com/facebookgo/grace/gracehttp"
	"go.uber.org/zap"
)

func main() {

	address, err := config.Config.String("address")
	if err != nil {
		mlog.Fatal("get config err", zap.Error(err))
	}

	s := &http.Server{
		Addr:              address,
		Handler:           api.Run(),
		ReadTimeout:       60 * time.Second,
		ReadHeaderTimeout: 60 * time.Second,
		IdleTimeout:       300 * time.Second,
		WriteTimeout:      20 * time.Second,
	}
	if err := gracehttp.Serve(s); err != nil {
		mlog.Fatal(err.Error())
	}
}
