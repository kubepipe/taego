package mhttp

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	"taego/lib/mlog"
)

type Client struct {
	http.Client
	host   string
	header http.Header
}

func NewDefaultClient(host string, header http.Header) *Client {
	timeout := 3 * time.Second

	return &Client{
		Client: http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				MaxIdleConns:          50,
				MaxIdleConnsPerHost:   50,
				IdleConnTimeout:       60 * time.Second,
				DisableCompression:    true,
				ResponseHeaderTimeout: timeout,
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					mlog.Debug("connectAddrList", "addr", addr)
					c, err := net.DialTimeout("tcp", addr, time.Second)
					if err != nil {
						mlog.Error("ConnectError", "addr", addr, "err", err)
						return nil, err
					}
					mlog.Debug("connectHost", "addr", addr)
					return c, nil
				},
			},
		},
		host:   host,
		header: header,
	}
}

// if header is nil, c.header will be used by default
func (c *Client) Get(ctx context.Context, path string, header http.Header,
) (int, []byte, error) {

	return c.call(ctx, http.MethodGet, path, header, nil)
}

// if header is nil, c.header will be used by default
func (c *Client) Post(ctx context.Context, path string, body []byte, header http.Header,
) (int, []byte, error) {

	return c.call(ctx, http.MethodPost, path, header, body)
}

// if header is nil, c.header will be used by default
func (c *Client) Put(ctx context.Context, path string, body []byte, header http.Header,
) (int, []byte, error) {

	return c.call(ctx, http.MethodPut, path, header, body)
}

// if header is nil, c.header will be used by default
func (c *Client) Delete(ctx context.Context, path string, body []byte, header http.Header,
) (int, []byte, error) {

	return c.call(ctx, http.MethodDelete, path, header, body)
}

// if header is nil, c.header will be used by default
func (c *Client) call(ctx context.Context, method, path string, header http.Header, body []byte,
) (int, []byte, error) {

	if c == nil {
		return 500, nil, errors.New("client is nil")
	}
	if c.header == nil {
		c.header = make(http.Header)
	}

	req := &http.Request{
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Method:     method,
		Host:       c.host,
		URL: &url.URL{
			Scheme: "http",
			Host:   c.host,
			Path:   path,
		},
		Header: c.header,
	}

	req.Header.Set("Host", c.host)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "*/*")

	if body != nil {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		req.Body = io.NopCloser(bytes.NewReader(body))
		req.ContentLength = int64(len(body))
	}

	if header != nil {
		for h, vs := range header {
			if len(vs) == 0 {
				continue
			}
			req.Header.Set(h, vs[0])
			for _, v := range vs[1:] {
				req.Header.Add(h, v)
			}
		}
	}

	resp, err := c.Do(req)
	if err != nil {
		return 500, nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	// TODO trace replace this

	return resp.StatusCode, b, err
}
