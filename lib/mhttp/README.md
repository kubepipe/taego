# mhttp

mhttp定义了一个Client结构体作为发起http连接的客户端：

```go
type Client struct {
	http.Client
	host   string
	header http.Header
}
```

每个host需要一个Client结构体，Client结构体包含一个header，作为访问host时默认携带的http请求头，当发起一个http请求时，会与自定义请求头合并发送给host，自定义请求头优先级高于默认请求头.

mhttp提供以下http方法：

## get

```go
(c *Client) Get(ctx context.Context, path string, header http.Header) (int, []byte, error)
```

## post

```go
(c *Client) Post(ctx context.Context, path string, body []byte, header http.Header) (int, []byte, error)
```

## put

```go
(c *Client) Put(ctx context.Context, path string, body []byte, header http.Header) (int, []byte, error)
```

## delete

```go
(c *Client) Delete(ctx context.Context, path string, body []byte, header http.Header) (int, []byte, error)
```

mhttp一般在service中使用，例如，访问baidu.com的demo如下：

1.创建一个.go文件为server/example/example.go

```go
package example

import (
	"context"
	"fmt"
	"net/http"

	"taego/lib/config"
	"taego/lib/mhttp"
)

var client *mhttp.Client

func init() {
	header := http.Header{}
	//header.Add("example-token-key", config.Config.UString("example.token", "example"))
	client = mhttp.NewDefaultClient(config.Config.UString("example.host", "baidu.com"), header)
}
```

定义一个变量client，并且在init函数中初始化，之后对baidu.com的访问全都基于这个初始化后的client

2.编写函数调用mhttp的Get方法：

```go
type (
	ReqExample struct{}
	ResExample []byte
)

func GetExampleData(ctx context.Context, req *ReqExample) (ResExample, error) {
	httpCode, resBody, err := client.Get(ctx, "/", nil)
	if err != nil {
		return nil, err
	}
	if httpCode != http.StatusOK {
		return nil, fmt.Errorf("http code %d", httpCode)
	}

	// or unmarshal
	res := resBody

	return res, nil
}
```

