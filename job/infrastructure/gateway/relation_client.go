package gateway

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/transport/tcp"
	"github.com/pkg/errors"
	"github.com/tinyhole/im/idl/mua/im/relation"
	"github.com/tinyhole/im/job/domain/gateway"
	"github.com/tinyhole/im/job/infrastructure/config"
)

type relationClient struct {
	svc relation.RelationService
}

func NewRelationClient(conf *config.BaseConfig) gateway.RelationClient {
	registry := etcd.NewRegistry(registry.Addrs(conf.ConfigCenterAddr))
	cli := client.NewClient(client.Registry(registry), client.Transport(tcp.NewTransport()))
	relationSvc := relation.NewRelationService("", cli)
	return &relationClient{
		svc: relationSvc,
	}
}

func (r *relationClient) GetGroupType(groupID int64) (int32, error) {
	req := &relation.GetGroupReq{
		GroupID: groupID,
	}
	rsp, err := r.svc.GetGroup(context.Background(), req)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return int32(rsp.Info.Type), nil
}


func (r *relationClient)ListGroupMember(groupID int64,page, pageSize int32)(rets []int64,total int32, err error){
	req := &relation.ListGroupMembersReq{
		GroupID:              groupID,
		Page:                 page,
		PageSize:             pageSize,
	}
	rsp, err := r.svc.ListGroupMembers(context.Background(), req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	total = rsp.Total
	rets = rsp.Uid
	return
}
