# Introduction

一个golang apiserver的脚手架, 在不引入复杂性（学习成本）的前提下保证功能性.

主要使用到的第三方开源工具:

* 路由: gin

* 日志: zap

# Feature

### trace模块

trace与zap log、gin完美配合, 使每一条由gin接收的请求在链路的关键点（如http、mysql等io调用）都有日志打印, 最后将trace.id返回给客户端.

使用curl模拟客户端请求：

```bash
➜  ~ curl 127.0.0.1:9091/api/v1/example | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   231  100   231    0     0  11584      0 --:--:-- --:--:-- --:--:-- 16500
{
  "errcode": 0,
  "success": true,
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

支持自定义error，返回给客户端errcode，用于特殊场景下的错误标识.

### 设计哲学

微服务场景下各模块职责单一，要避免单个模块过度分层导致的开发、维护成本上升，因此taego从上到下依次为：

* 路由层 api：负责接口定义
* 逻辑层 controller：负责主要业务逻辑
* 调用层 dao/service：中间件或依赖服务的调用

# Directory

tree -I vendor
```
.
├── Makefile
├── README.md
├── api
│   ├── router.go
│   └── server.go
├── bin
│   └── server
├── controller
│   ├── auth.go
│   ├── common.go
│   └── example.go
├── dao
├── etc
│   └── config.yaml
├── go.mod
├── go.sum
├── lib
│   ├── config
│   │   └── config.go
│   ├── k8s
│   │   └── clients.go
│   ├── mhttp
│   │   └── http.go
│   ├── mlog
│   │   └── log.go
│   ├── trace
│   │   └── trace.go
│   ├── util
│   │   └── util.go
│   └── watcher
│       └── watcher.go
├── license
├── main.go
├── mconst
│   ├── const.go
│   ├── errors.go
│   ├── request.go
│   └── response.go
└── service
    └── example
        └── example.go

16 directories, 25 files
```

* main.go 入口函数
* api server启动、路由
* controller 主要业务逻辑
* bin 编译后的可执行文件
* etc 配置文件
* mconst 全局常量、变量定义
* dao 数据增删改查
* lib 对基础依赖的封装,如http、mysql、k8s、log
* service 第三方依赖服务调用

# Roadmap 

https://github.com/orgs/kubepipe/projects/1

# license

MIT

