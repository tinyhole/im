package gateway

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/transport/tcp"
	"github.com/pkg/errors"
	"github.com/tinyhole/im/idl/mua/im/ap"
	"github.com/tinyhole/im/job/domain/gateway"
	"github.com/tinyhole/im/job/infrastructure/config"
)

type apClient struct {
	apSvc ap.APService
}

func NewApClient(conf *config.BaseConfig) gateway.ApClient {
	registry := etcd.NewRegistry(registry.Addrs(conf.ConfigCenterAddr))
	cli := client.NewClient(client.Registry(registry), client.Transport(tcp.NewTransport()))
	apSvc := ap.NewAPService("", cli)

	return &apClient{
		apSvc: apSvc,
	}
}

func (a *apClient) Unicast(apID int32, fid int64, data []byte) (err error) {
	req := &ap.UnicastReq{
		Fid:      fid,
		SrvName:  "mua.im.job",
		Endpoint: "Job.PushMsg",
		Data:     data,
	}
	_, err = a.apSvc.Unicast(context.Background(), req,
		client.WithSelectOption(selector.WithFilter(a.FilterID(fmt.Sprintf("mua.im.ap-%d", apID)))))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (a *apClient) Broadcast(apID int32, fids []int64, data []byte) (err error) {
	req := &ap.BroadcastReq{
		Fids:     fids,
		SrvName:  "mua.im.job",
		Endpoint: "Job.PushMsg",
		Data:     data,
	}
	_, err = a.apSvc.Broadcast(context.Background(), req,
		client.WithSelectOption(selector.WithFilter(a.FilterID(fmt.Sprintf("mua.im.ap-%d", apID)))))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (a *apClient) FilterID(id string) selector.Filter {
	return func(services []*registry.Service) (rets []*registry.Service) {
		for _, itr := range services {
			svc := registry.Service{
				Name:      itr.Name,
				Version:   itr.Version,
				Metadata:  itr.Metadata,
				Endpoints: itr.Endpoints,
			}
			for _, node := range itr.Nodes {
				//fmt.Printf("node [%v,dst[%v]",node.Id, id)
				if node.Id == id {
					svc.Nodes = append(svc.Nodes, node)
					rets = append(rets, &svc)
				}
			}
		}
		return
	}
}
