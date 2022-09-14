# Introduction

一个golang apiserver的脚手架

# Feature

* 轻量
* Gopher 

# Director

tree -I vendor
```
.
├── LICENSE
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
│   ├── const.go
│   ├── errors.go
│   ├── request.go
│   └── response.go
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
├── main.go
└── service
    └── example
        └── example.go

15 directories, 25 files
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

# Some else

welcome join!


