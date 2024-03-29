package server

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/transport/tcp"
	"github.com/pkg/errors"
	"github.com/tinyhole/im/ap/logger"
	"github.com/tinyhole/im/ap/tcpserver/bucket"
	"github.com/tinyhole/im/ap/tcpserver/protocol/pack"
	"github.com/tinyhole/im/ap/tcpserver/transport"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type apServer struct {
	opts      *Options
	transport transport.Transport
	rpcClient client.Client
	ctx       context.Context
	cancelFn  context.CancelFunc
	listener  transport.Listener
	log       logger.Logger
}

func NewAPServer(log logger.Logger, opts ...Option) Server {
	a := &apServer{
		log: log,
	}
	a.opts = newOptions()
	for _, o := range opts {
		o(a.opts)
	}
	a.transport = transport.NewTcpTransport(log)
	a.ctx, a.cancelFn = context.WithCancel(context.Background())
	a.rpcClient = client.NewClient(client.Registry(etcd.NewRegistry()), client.Transport(tcp.NewTransport()))
	return a
}

func (a *apServer) Start() error {
	listener, err := a.transport.Listen(a.opts.Addr)

	if err != nil {
		return errors.Wrap(err, "transport.Listen failed")
	}

	listener.Accept(a.DestroyCon, a.ProcessMsg)
	a.listener = listener
	return nil
}

func (a *apServer) Stop() error {
	a.listener.Close()
	a.cancelFn()
	return nil
}

func (a *apServer) DestroyCon(socket transport.Socket) {
	fID := a.GenerateFid(socket.ID())
	bucket.DefaultSocketBucket.Remove(fID)
}

func (a *apServer) ConHeartbeat(socket transport.Socket) {
	fid := a.GenerateFid(socket.ID())
	if socket.GetAuthState() == true {
		a.opts.sessClient.Ping(socket.GetUID(), fid, int32(a.opts.SrvID))
	}
}

func (a *apServer) authSocket(socket transport.Socket, header *pack.Header) (err error) {
	//1.身份信息
	uid := header.Auth.Uid
	token := header.Auth.Token
	//2.认证
	if a.opts.authClient != nil {
		fid := a.GenerateFid(socket.ID())
		err = a.opts.authClient.SignIn(uid, fid, int32(a.opts.SrvID), int32(header.Device.Type),
			"mua.im.ap", token, socket.Remote())
		if err == nil {
			socket.UpdateAuthState(true)
			socket.SetUID(uid)
			bucket.DefaultSocketBucket.Add(fid, socket)
		} else {
			return ErrAuthFailed
		}
	}
	return nil
}

const (
	CallTypeUnknown = 0
	Sync            = 1
	Async           = 2
	Push            = 3
)

func (a *apServer) ProcessMsg(socket transport.Socket, reqPack *pack.ApPackage) {
	defer func() {
		if r := recover(); r != nil {
			a.log.Errorf("process msg failed [%v]", r)
		}
	}()

	var (
		err    error
		reqTmp *Message
		rspTmp *Message
	)
	if socket.GetAuthState() == false {
		err = a.authSocket(socket, reqPack.Header)
		if err != nil {
			a.log.Warnf("socket auth failed [%v]", socket.Remote())
			socket.Close()
		}

	}
	if reqPack.Header.Request != nil {
		if reqPack.Header.Request.ServiceName == "mua.im.ap" &&
			reqPack.Header.Request.Endpoint == "AP.Ping" {
			a.ConHeartbeat(socket)
			return
		}

		if reqPack.Body != nil {
			reqTmp = NewMessage(reqPack.Body)
		}
		req := a.rpcClient.NewRequest(reqPack.Header.Request.ServiceName, reqPack.Header.Request.Endpoint, reqTmp)
		rspPack := &pack.ApPackage{
			Header: &pack.Header{
				Response: &pack.ResponseMeta{
					ErrCode: 0,
					ErrText: "",
				},
				Seq: reqPack.Header.Seq,
			},
		}
		rspTmp = NewMessage([]byte{})
		if err = a.rpcClient.Call(a.ctx, req, rspTmp); err != nil {
			statu, flag := status.FromError(err)
			if flag {
				rspPack.Header.Response.ErrCode = int32(statu.Code())
				rspPack.Header.Response.ErrText = statu.Message()
			} else {
				rspPack.Header.Response.ErrCode = int32(codes.Unknown)
				rspPack.Header.Response.ErrText = err.Error()
			}

		} else {
			rspPack.Header.Response.ErrCode = OK
			rspPack.Body = rspTmp.data
		}
		if reqPack.Header.Request.CallType == Sync {
			socket.Send(rspPack)
		}
	}

	return
}

func (a apServer) GenerateFid(id uint32) int64 {
	base := uint64(a.opts.SrvID) << 32
	return int64(base + uint64(id))
}
