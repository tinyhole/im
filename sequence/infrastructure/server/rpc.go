package server

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/transport/tcp"
	"github.com/tinyhole/im/sequence/infrastructure/config"
)

func NewRPCServer(conf *config.BaseConfig) micro.Service {
	svc := micro.NewService(micro.Transport(tcp.NewTransport()), micro.Name("mua.im.sequence"),
		micro.Registry(etcd.NewRegistry(registry.Addrs(conf.RegistryCenterAddr))))
	return svc
}
