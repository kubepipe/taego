package mhttp

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestNewDefaultClient(t *testing.T) {
	type args struct {
		host   string
		header http.Header
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := NewDefaultClient(tt.args.host, tt.args.header); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewDefaultClient() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestClient_Get(t *testing.T) {
	type args struct {
		ctx    context.Context
		path   string
		header http.Header
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    int
		want1   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, got1, err := tt.c.Get(tt.args.ctx, tt.args.path, tt.args.header)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Client.Get() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Client.Get() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. Client.Get() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestClient_Post(t *testing.T) {
	type args struct {
		ctx    context.Context
		path   string
		body   []byte
		header http.Header
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    int
		want1   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, got1, err := tt.c.Post(tt.args.ctx, tt.args.path, tt.args.body, tt.args.header)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Client.Post() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Client.Post() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. Client.Post() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestClient_Put(t *testing.T) {
	type args struct {
		ctx    context.Context
		path   string
		body   []byte
		header http.Header
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    int
		want1   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, got1, err := tt.c.Put(tt.args.ctx, tt.args.path, tt.args.body, tt.args.header)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Client.Put() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Client.Put() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. Client.Put() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestClient_Delete(t *testing.T) {
	type args struct {
		ctx    context.Context
		path   string
		body   []byte
		header http.Header
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    int
		want1   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, got1, err := tt.c.Delete(tt.args.ctx, tt.args.path, tt.args.body, tt.args.header)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Client.Delete() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Client.Delete() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. Client.Delete() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestClient_call(t *testing.T) {
	type args struct {
		ctx    context.Context
		method string
		path   string
		header http.Header
		body   []byte
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    int
		want1   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, got1, err := tt.c.call(tt.args.ctx, tt.args.method, tt.args.path, tt.args.header, tt.args.body)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Client.call() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Client.call() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. Client.call() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}
