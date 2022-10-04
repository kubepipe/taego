# Introduction

taego【泰戈】,一个轻量的golang apiserver脚手架, 在不引入复杂性（学习成本）的前提下保证功能性.

主要使用到的第三方开源工具:

* 路由: gin

* 日志: zap

## 结构

taego从上到下依次为：

* 路由层 api：负责接口定义
* 逻辑层 controller：负责主要业务逻辑，lib库封装的各模块，以及dao/service通常在controller中使用
* 调用层 dao/service：中间件或依赖服务的调用

# Features

### trace模块

api server的日志不加trace,不能跟请求对应起来的话,是没有意义的.

trace配合zap、gin, 使每一条由gin接收的请求在链路的关键点（如http、mysql等io调用）都有日志打印, 最后将trace.id返回给客户端.

在controller层的用法:
```
GetTrace(c).Log("some other things to do")
```

日志示例如下：

```
{"level":"info","ts":1664612799.451587,"caller":"mtrace/trace.go:60","msg":"step done","queryNum":2,"trace":1562664607,"traceName":"select name from user limit 10","totalTime":"452.25µs"}
{"level":"info","ts":1664612799.502675,"caller":"mtrace/trace.go:60","msg":"step done","trace":1562664607,"traceName":"GET-baidu.com/","totalTime":"50.958458ms"}
{"level":"info","ts":1664612799.502698,"caller":"mtrace/trace.go:49","msg":"some other things to do","trace":1562664607,"traceName":"GET-127.0.0.1:9091/api/v1/example"}
{"level":"info","ts":1664612799.502883,"caller":"mtrace/trace.go:60","msg":"step done","trace":1562664607,"traceName":"GET-127.0.0.1:9091/api/v1/example","totalTime":"51.758833ms"}
```

### msql

msql模块是对database/sql的封装.

sql还是orm?

sql是操作数据库的规范，但orm不是，每个orm框架都有自己的规范;

orm通过将高级语言翻译成sql，提高开发效率，但同时引入性能损耗.

taego主张完全由原生sql控制数据库的增删改查以及索引优化等，另外考虑到使用golang的原生database/sql包在执行批量查询时，代码过于繁琐，因此封装msql模块旨在提高开发效率，兼顾执行效率，又不引入学习成本.

使用示例：

首先初始化数据库连接：

```
var user msql.SQL

func init() {
	user = msql.NewSQL("user:password@/dbname")
}
```

定义一个struct表示表结构，其中每个字段的db tag表示对应表的字段名：

```
type User struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
```

select 使用SQL.Query方法，将结果存入目标变量中：

```
us := []*User{}
_ = user.Query(ctx, "select * from user limit 10").Scan(&us)
```

update / delete / insert 使用SQL.Exec方法：

```
id := 1
if _, err := user.Exec(ctx, "delete from user where id=?", id); err != nil {
	return err
}
```

### context.Context

gin结构体封装的http.Request包含一个Context，可用于客户端连接关闭时的通知.

taego使用request.Context()生成一个span context，贯穿一个请求的整个生命周期.

trace,user等元数据存放在span context中贯穿整个链路，当客户端请求关闭时，请求创建的goroutine都会得到通知，通常作为函数第一个参数.

# Document

[快速开始](docs/quick-start.md)

[数据库操作 lib/msql](lib/msql/README.md)

[trace模块 lib/mtrace](lib/mtrace/README.md)

[自定义error lib/merrors](lib/merrors/README.md)

[http模块 lib/mhttp](lib/mhttp/README.md)

[控制器模块 controller](controller/README.md)

# License

MIT

另外如果觉得不错麻烦给个🌟，如果想吐槽可以在github提issue，中英文皆可~
