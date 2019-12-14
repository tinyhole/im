package rpc

import (
	"context"
	"github.com/tinyhole/im/ap/tcpserver/client"
	"github.com/tinyhole/im/idl/mua/im/ap"
)

type Handler struct {
	tcpClient client.Client
}

func NewHandler(cli client.Client) *Handler {
	return &Handler{tcpClient: cli}
}

func(h *Handler)Unicast(ctx context.Context, req *ap.UnicastReq, rsp *ap.UnicastRsp)(err error){
	h.tcpClient.Unicast(req.Fid, req.SrvName, req.Endpoint, req.Data)
	return nil
}

func (h *Handler)Broadcast(ctx context.Context, req *ap.BroadcastReq, rsp *ap.BroadcastRsp)(err error){
	h.tcpClient.Broadcast(req.Fids, req.SrvName, req.Endpoint, req.Data)
	return nil
}
