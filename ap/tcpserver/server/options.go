package server

import (
	"github.com/tinyhole/im/ap/tcpserver/gateway"
)

type Options struct {
	Addr       string
	SrvID      uint32
	sessClient gateway.SessionClient
	authClient gateway.AuthClient
}

type Option func(o *Options)

func newOptions() *Options {
	return &Options{
		Addr: ":8080",
	}
}

func WithLocalAddr(addr string) Option {
	return func(o *Options) {
		o.Addr = addr
	}
}

func WithSrvID(id uint32) Option {
	return func(o *Options) {
		o.SrvID = id
	}
}

func WithSessionClient(client gateway.SessionClient) Option {
	return func(o *Options) {
		o.sessClient = client
	}
}

func WithAuthClient(client gateway.AuthClient) Option {
	return func(o *Options) {
		o.authClient = client
	}
}
