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

假设接口路径为`/demo`，get方法，接口返回一个简单字符串`demo`

首先在路由层定义http方法、接口路径以及处理函数

在api/router.go中增加如下代码：

```go
package api

import (
	ctl "taego/controller"

	"github.com/gin-gonic/gin"
)

func setRoute(e *gin.Engine) {
	e.Any("/", ctl.Ok)
	e.GET("/health", ctl.Health)

	e.GET("/demo", ctl.Demo)

	v1 := e.Group("/api/v1", ctl.Auth)

	v1.GET("/example", ctl.Example)
}
```

其中`e.GET("/demo", ctl.Demo)`是新增的demo接口

接下来创建一个controller/demo.go文件，并在文件中加入以下代码：

```go
package controller

import "github.com/gin-gonic/gin"

func Demo(c *gin.Context) {
	success(c, "demo")
}
```

到这里demo接口就写完了,然后访问一下看看效果：

运行`make run`来启动api server

使用浏览器访问`http://127.0.0.1/demo`

返回如下结果：

```
{
  "errcode": 0,
  "trace": {
    "id": 657218595,
    "sourceIp": "127.0.0.1",
    "serverIp": "192.168.31.29"
  },
  "data": "demo"
}
```

响应中的data字段就是controller/demo.go返回的内容，而其余字段各含义如下：

* errcode: 错误码，返回0表示正常状态。用于特殊场景下返回指定错误，例如当errcode=10000时表示未登录，需要跳转到登录页面
* trace: 当前请求的唯一标识，根据trace.id查询日志，可用于debug，统计请求耗时，标记某个指定步骤的耗时，串联一次请求输出的日志
* data: controller方法的输出会放在data中返回给客户端



