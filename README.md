# Introduction

taego【泰戈】,一个轻量的golang apiserver脚手架, 在不引入复杂性（学习成本）的前提下保证功能性.

主要使用到的第三方开源工具:

* 路由: gin

* 日志: zap

## 结构

taego从上到下依次为：

* 路由层 api：负责接口定义
* 逻辑层 controller：负责主要业务逻辑
* 调用层 dao/service：中间件或依赖服务的调用

# Features

### trace模块

api server的日志不加trace,不能跟请求对应起来的话,是没有意义的.

trace配合zap、gin, 使每一条由gin接收的请求在链路的关键点（如http、mysql等io调用）都有日志打印, 最后将trace.id返回给客户端.

在controller层的用法:
```
GetTrace(c).Log("this is my log")
```

使用curl模拟客户端请求：

```bash
➜  ~ curl 127.0.0.1:9091/api/v1/example | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   231  100   231    0     0  11584      0 --:--:-- --:--:-- --:--:-- 16500
{
  "errcode": 0,
  "trace": {
    "id": 590419775,
    "sourceIp": "127.0.0.1",
    "serverIp": "192.168.31.29"
  },
  "data": "<html>\n<meta http-equiv=\"refresh\" content=\"0;url=http://www.baidu.com/\">\n</html>\n"
}
```

日志如下：

```
{"level":"info","ts":1663552664.8669891,"caller":"mtrace/trace.go:50","msg":"i am doing some curd","trace":590419775,"traceName":"demohandle"}
{"level":"info","ts":1663552664.8670452,"caller":"mtrace/trace.go:56","msg":"step done","trace":590419775,"traceName":"demohandle","totalTime":"55.958µs"}
{"level":"info","ts":1663552664.8808,"caller":"mtrace/trace.go:56","msg":"step done","trace":590419775,"traceName":"GET-baidu.com/","totalTime":"13.737666ms"}
{"level":"info","ts":1663552664.880825,"caller":"mtrace/trace.go:50","msg":"some other things to do","trace":590419775,"traceName":"GET-127.0.0.1:9091/api/v1/example"}
{"level":"info","ts":1663552664.880856,"caller":"mtrace/trace.go:56","msg":"step done","trace":590419775,"traceName":"GET-127.0.0.1:9091/api/v1/example","totalTime":"13.887958ms"}
```

### merrors模块

自定义error，返回给客户端errcode，用于特殊场景下的错误标识.

lib/merrors中定义：

```go
const (
	ERROR_UNAUTHORIZED Code = iota + 10000
	ERROR_UNHEALTHY
	// TODO add new error code here

)

var errmap = map[string]Code{
	// TODO add new error descriptions here

	"unauthorized": ERROR_UNAUTHORIZED,
	"unhealthy":    ERROR_UNHEALTHY,
}
```

controller中使用:

```go
fail(c, merrors.Get(merrors.ERROR_UNAUTHORIZED))
```

响应预览：

```bash
{
  "errcode": 10000,
  "message": "unauthorized",
  "trace": {
    "id": 2005758541,
    "sourceIp": "127.0.0.1",
    "serverIp": "192.168.31.29"
  }
}
```

### context.Context

gin结构体封装的http.Request包含一个Context，可用于客户端连接关闭时的通知.

taego使用request.Context()生成一个span context，贯穿一个请求的整个生命周期.

trace,user等元数据存放在span context中贯穿整个链路，当客户端请求关闭时，请求创建的goroutine都会得到通知，通常作为函数第一个参数.

### mmysql

mmysql模块是对mysql sdk的封装，使用go-sql-driver驱动.

为什么不用orm，主要考虑到学习成本以及性能损耗，sql是操作数据库的规范，但orm不是，每个orm框架都有自己的规范，且orm大多使用反射，这会带来性能上的损耗.

taego主张使用原生sql，完全由原生sql控制数据库的增删改查以及索引优化等，另外学习zap通过指定数据类型的方式代替反射.

# Document

[快速开始](docs/quick-start.md)

# Directory

```
➜  taego git:(master) ✗ tree -I vendor -L 2
.
├── Makefile
├── README.md
├── api
│   ├── README.md
│   ├── router.go
│   ├── router_test.go
│   ├── server.go
│   └── server_test.go
├── bin
│   └── server
├── controller
│   ├── README.md
│   ├── auth.go
│   ├── auth_test.go
│   ├── common.go
│   ├── common_test.go
│   ├── example.go
│   ├── example_test.go
│   ├── health.go
│   └── health_test.go
├── dao
│   └── README.md
├── docs
│   └── quick-start.md
├── etc
│   ├── README.md
│   └── config.yaml
├── go.mod
├── go.sum
├── lib
│   ├── README.md
│   ├── config
│   ├── k8s
│   ├── merrors
│   ├── mhttp
│   ├── mlog
│   ├── mtrace
│   ├── util
│   └── watcher
├── license
├── main.go
├── mconst
│   ├── README.md
│   ├── const.go
│   ├── request.go
│   ├── response.go
│   └── user.go
└── service
    ├── README.md
    └── example

18 directories, 32 files
```

* main.go: 入口函数
* api: server启动、路由
* bin: 编译后的可执行文件
* controller: 主要业务逻辑
* dao: 数据访问
* etc: 配置文件
* lib: 对基础依赖的封装,如http、mysql、k8s、log、trace
* mconst: 全局常量、变量定义
* service 第三方依赖服务调用

# Roadmap 

https://github.com/orgs/kubepipe/projects/1

# License

MIT

