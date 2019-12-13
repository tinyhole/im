package gateway

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/transport/tcp"
	"github.com/pkg/errors"
	"github.com/tinyhole/im/idl/mua/im/logic"
	"github.com/tinyhole/im/job/domain/gateway"
	"github.com/tinyhole/im/job/domain/valueobj"
	"github.com/tinyhole/im/job/infrastructure/config"
	"github.com/tinyhole/im/job/interfaces/objconv"
)

type sessionClient struct {
	sessionSvc logic.LogicService
}

func NewSessionClient(conf *config.BaseConfig) gateway.SessionClient {
	cli := client.NewClient(client.Registry(etcd.NewRegistry(registry.Addrs(conf.ConfigCenterAddr))),
		client.Transport(tcp.NewTransport()))
	svc := logic.NewLogicService("", cli)

	return &sessionClient{sessionSvc: svc}
}

func (s *sessionClient) ListSessionInfo(uid int64) ([]*valueobj.SessionInfo, error) {
	req := &logic.ListSessionInfoReq{
		Uid: uid,
	}
	rsp, err := s.sessionSvc.ListSessionInfo(context.Background(), req)

	if err != nil {
		return nil, errors.WithStack(err)
	}
	return objconv.SessionInfoConv.SliceDTO2DO(rsp.Infos), nil
}

func (s *sessionClient) BatchListSessionInfo(uids []int64) (ret map[int64][]*valueobj.SessionInfo, err error) {
	req := &logic.BatchListSessionInfoReq{
		Uids: uids,
	}
	rsp, err := s.sessionSvc.BatchListSessionInfo(context.Background(), req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rets := objconv.SessionInfoConv.SliceDTO2DO(rsp.Infos)
	retMap := make(map[int64][]*valueobj.SessionInfo)
	for _, itr := range rets {
		if v, ok := retMap[itr.Uid]; ok {
			v = append(v, itr)
		} else {
			retMap[itr.Uid] = []*valueobj.SessionInfo{itr}
		}
	}

	return retMap, nil
}
