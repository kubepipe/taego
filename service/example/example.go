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

type (
	ReqExample struct{}
	ResExample []byte
)

func GetExampleData(ctx context.Context, req *ReqExample) (ResExample, error) {
	httpCode, resBody, err := client.Get(ctx, "/", client.DefaultHeader())
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
