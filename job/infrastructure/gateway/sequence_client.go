package gateway

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/transport/tcp"
	"github.com/pkg/errors"
	"github.com/tinyhole/im/idl/mua/im/sequence"
	"github.com/tinyhole/im/job/domain/gateway"
	"github.com/tinyhole/im/job/infrastructure/config"
)

type sequenceClient struct {
	svc sequence.SequenceService
}

func NewSequenceClient(conf *config.BaseConfig) gateway.SequenceClient {
	registry := etcd.NewRegistry(registry.Addrs(conf.ConfigCenterAddr))
	cli := client.NewClient(client.Registry(registry), client.Transport(tcp.NewTransport()))
	svc := sequence.NewSequenceService("", cli)
	return &sequenceClient{
		svc: svc,
	}
}

func (s *sequenceClient) GetPrivateSeq(inboxID string) (int64, error) {
	req := &sequence.GetAutoIncrIDReq{
		Key: inboxID,
	}

	rsp, err := s.svc.GetAutoIncrID(context.Background(), req)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return rsp.Id, err
}

func (s *sequenceClient) GetGroupSeq(inboxID string) (int64, error) {
	req := &sequence.GetAutoIncrIDReq{
		Key: inboxID,
	}

	rsp, err := s.svc.GetAutoIncrID(context.Background(), req)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return rsp.Id, err
}
