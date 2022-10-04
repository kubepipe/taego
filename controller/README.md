# controller

controller是taego处理业务逻辑的地方，查询数据库、调用依赖服务获取到的数据，在controller中处理后通过controller提供的公共方法返回给调用方.

controller提供的用来给调用方返回数据的方法有3个：

1.成功情况使用success返回数据：

```go
success(c *gin.Context, obj any) 
```

2.处理失败，但使用http 200状态码返回数据，通过errcode标识错误原因：

```go
fail(c *gin.Context, merr merrors.Merr) 
```

3.处理失败，使用指定的http状态码返回数据：

```go
failNot200(c *gin.Context, httpcode int, merr merrors.Merr)
```



另外提供getSpan方法获取context.Context：

```go
getSpan(c *gin.Context) context.Context
```

getTrace方法获取Context中的trace数据：

```go
getTrace(c *gin.Context) mtrace.Trace
```

