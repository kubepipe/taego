.PHONY: all build run

all: build

pre:
	go mod vendor
	mkdir -p bin

clean:
	rm -rf vendor
	rm -rf bin
	rm -rf main_test.go

build:clean pre
	GOPROXY=https://goproxy.cn,direct GO111MODULE=on CGO_ENABLED=0 go build -o bin/server main.go

run:build
	# so why not use ./bin/server?
	# becuase go test. The etc directory must be the parent directory
	cd bin && MODE=release ./server

test:
	go test ./... -cover -v

generate-test:build
	gotests -w -all lib/config
	gotests -w -all lib/k8s
	gotests -w -all lib/merrors
	gotests -w -all lib/mhttp
	gotests -w -all lib/mlog
	gotests -w -all lib/mtest
	gotests -w -all lib/mtrace
	gotests -w -all lib/util
	gotests -w -all lib/watcher
	\
	gotests -w -all api/
	\
	gotests -w -all controller
