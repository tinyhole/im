package gateway

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/transport/tcp"
	"github.com/tinyhole/im/ap/config"
	"github.com/tinyhole/im/idl/mua/im/logic"
)

type AuthClient interface {
	SignIn(uid, fid int64, apID, deviceType int32, srvName, token, remoteIP string) error
}

type authClient struct {
	svc logic.LogicService
}

func NewAuthClient(conf *config.BaseConfig) AuthClient {
	registry := etcd.NewRegistry(registry.Addrs(conf.ConfigCenterAddr))
	cli := client.NewClient(client.Registry(registry), client.Transport(tcp.NewTransport()))
	svc := logic.NewLogicService("", cli)
	fmt.Println("new auth client")
	return &authClient{
		svc: svc,
	}
}

func (a *authClient) SignIn(uid, fid int64, apID, deviceType int32, srvName, token, remoteIP string) error {
	req := &logic.SignInReq{
		Uid:        uid,
		Token:      token,
		ApID:       apID,
		ApFid:      fid,
		RemoteIP:   remoteIP,
		DeviceType: deviceType,
		SrvName:    srvName,
	}
	_, err := a.svc.SignIn(context.Background(), req)
	return err
}

type SessionClient interface {
	Ping(uid int64, fid int64, srvID int32)
}

type sessionClient struct {
	svc logic.LogicService
}

func NewSessionClient(conf *config.BaseConfig) SessionClient {
	registry := etcd.NewRegistry(registry.Addrs(conf.ConfigCenterAddr))
	cli := client.NewClient(client.Registry(registry), client.Transport(tcp.NewTransport()))
	svc := logic.NewLogicService("", cli)
	return &sessionClient{svc: svc}
}

func (s *sessionClient) Ping(uid, fid int64, srvID int32) {
	req := &logic.PingReq{
		Uid:   uid,
		ApID:  srvID,
		ApFid: fid,
	}
	s.svc.Ping(context.Background(), req)
}
