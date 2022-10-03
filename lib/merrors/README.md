# merrors

taego自定义一个Merr类型用于向调用方返回error code，Merr接口定义如下：

```go
type Merr interface {
	Error() string
	Code() int
}
```

相比于golang的error类型，Merr多了一个Code方法，目的在于标识特定error，调用方可根据特定error code执行特定操作.

如何增加一个Merr？

lib/merrors/merr.go:

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

以ERROR_UNAUTHORIZED为例，ERROR_UNAUTHORIZED本身为code类型，值为10000，其后的每一行code都自动加1，再给ERROR_UNAUTHORIZED定义一个描述作为返回给调用方的message: "unauthorized".

在controller中使用ERROR_UNAUTHORIZED:

```go
unauthorized := true
if unauthorized {
	fail(c, merrors.Get(merrors.ERROR_UNAUTHORIZED))
	return
}
```

接口返回如下:

```json
{
  "errcode": 10000,
  "message": "unauthorized",
  "trace": {
    "id": 1254329389,
    "sourceIp": "127.0.0.1",
    "serverIp": "192.168.31.29"
  }
}
```

