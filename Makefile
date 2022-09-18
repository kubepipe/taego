.PHONY: all build run

all: build

pre:
	go mod vendor
	mkdir -p bin

clean:
	rm -rf vendor
	rm -rf bin

build:clean pre
	GO111MODULE=on CGO_ENABLED=0 go build -o bin/server main.go

run:build
	# so why not use ./bin/server?
	# becuase go test. The etc directory must be the parent directory
	MODE=release cd bin && ./server

test:
	go test ./... -cover -v
