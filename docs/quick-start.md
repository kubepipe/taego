# Quick start

本节目标在于利用taego写一个demo接口，并可通过curl或浏览器访问。

环境：

* golang >= 1.18

## 启动

从github下载代码后，通过自带的makefile启动server：

mac或linux运行：

```go
make run
```

或可通过`GOPROXY=https://goproxy.cn,direct GO111MODULE=on CGO_ENABLED=0 go build -o bin/server main.go`编译，然后通过`MODE=release ./bin/server`运行。

成功后的运行结果：

```bash
➜  taego git:(master) ✗ make run
rm -rf vendor
rm -rf bin
rm -rf main_test.go
go mod vendor
mkdir -p bin
GOPROXY=https://goproxy.cn,direct GO111MODULE=on CGO_ENABLED=0 go build -o bin/server main.go
# so why not use ./bin/server?
# becuase go test. The etc directory must be the parent directory
cd bin && MODE=release ./server
```

taego默认监听端口为9091，可通过etc/config.yaml修改address字段。

通过浏览器访问`http://127.0.0.1:9091/` 页面返回如下结果:

```
{"errcode":0,"success":true,"trace":{"id":298074068,"sourceIp":"127.0.0.1","serverIp":"192.168.31.29"},"data":"ok"}
```

同时，终端输出：

```
{"level":"info","ts":1663600066.043964,"caller":"mtrace/trace.go:55","msg":"step done","trace":298074068,"traceName":"GET-127.0.0.1:9091/","totalTime":"630.125µs"}
```

## 写一个demo

假设接口路径为`/demo`,接口返回一个简单字符串`demo`

首先在路由层定义http方法、接口路径以及处理函数

