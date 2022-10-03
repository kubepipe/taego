# mtrace

trace可用于标记某一段业务逻辑的开始结束时间，并且附加一些额外信息与traceid一起打印到log中.

taego的trace通过以下方法实现：

```go
type Trace interface {
	// return traceid
	GetTraceId() int32

	// You can count the total time spent here, print logs, and upload monitoring metrics
	Done(args ...zap.Field)

	// generate child trace
	SubTrace(string) Trace

	// log with traceid
	Log(string, ...zap.Field)
}
```

如何在controller层使用标记一段业务代码的开始结束时间，并且附加一些信息打印到日志中：

```go
func demo(c *gin.Context) {
  // ...
	trace1 := mtrace.SubTrace(GetSpan(c), "trace1 name")
	// do some thing ...
  // optional
  trace1.Log("xxx")
	trace1.Done(zap.String("key", "value"))
	// ...
}
```

首先使用mtrace.SubTrace方法生成一个子trace对象，并且传入一个name用于标识这个子trace，最后使用trace.Done方法结束子trace.

Done方法内部会将传入的参数，与traceid一起打印到日志中，也可以使用trace.Log方法显式地打印日志.

此时访问该接口，日志中会出现以下日志：

```
{"level":"info","ts":1664788275.1827788,"caller":"mtrace/trace.go:49","msg":"xxx","trace":1158602154,"traceName":"trace1 name"}
{"level":"info","ts":1664788275.1830258,"caller":"mtrace/trace.go:60","msg":"step done","key":"value","trace":1158602154,"traceName":"trace1 name","totalTime":"248.125µs"}
```

### 实现原理：

taego利用gin路由框架，在请求开始时生成一个trace对象，存放在gin.Context中.

api/server.go:

```go
	e := gin.New()
	e.Use(ctl.SetSpan)
```

SetSpan在controller层的公共方法中:

controller/common.go:

```go
func SetSpan(c *gin.Context) {
	span, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	// trace
	name := fmt.Sprintf("%s-%s%s",
		c.Request.Method, c.Request.Host, c.Request.RequestURI)
	trace := mtrace.New(name)
	defer trace.Done()

	span = mtrace.ContextWithTrace(span, trace)
	c.Set(spanKey, span)

	c.Next()
}
```

可以看到taego利用c.Request.Context生成一个Context，并且在请求结束后通过cancel方法通过Context通知所有goroutine.

首先通过mtrace提供的New方法生成一个trace对象，trace的name为"http方法名-host+uri"，在请求结束后调用Done方法标记trace结束，最后将生成的trace对象存放在gin.Context中.

taego的数据库访问msql包、http调用mhttp包都默认加上了trace标识，以mhttp为例：

```go
	trace := mtrace.SubTrace(ctx, fmt.Sprintf("%s-%s%s", method, c.host, path))
	defer trace.Done()
```

