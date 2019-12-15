package client

import (
	"github.com/pkg/errors"
	"github.com/tinyhole/im/ap/tcpserver/bucket"
	"github.com/tinyhole/im/ap/tcpserver/protocol/pack"
)

type Client interface {
	Unicast(fid int64, service, endpoint string, data []byte) error
	Broadcast(fids []int64, service, endpoint string, data []byte) error
}

type client struct{}

var (
	APClient = &client{}
)

func (c *client) Unicast(fid int64, service, endpoint string, data []byte) error {
	socket, err := bucket.DefaultSocketBucket.Get(fid)
	if err != nil {
		return errors.Wrapf(err, "DefaultSocketBucket.Get(%d)", fid)
	}

	req := &pack.ApPackage{
		Header: &pack.Header{
			Request: &pack.RequestMeta{
				ServiceName: service,
				Endpoint:    endpoint,
				CallType:    pack.CallType_Push,
			},
			Seq: 0,
		},
		Body: data,
	}
	return socket.Send(req)
}

func (c *client) Broadcast(fids []int64, service, endpoint string, data []byte) error {
	req := &pack.ApPackage{
		Header: &pack.Header{
			Request: &pack.RequestMeta{
				ServiceName: service,
				Endpoint:    endpoint,
				CallType:    pack.CallType_Push,
			},
			Seq: 0,
		},
		Body: data,
	}
	for _, fid := range fids {
		socket, err := bucket.DefaultSocketBucket.Get(fid)
		if err != nil {
			continue
		}
		socket.Send(req)
	}

	return nil
}
