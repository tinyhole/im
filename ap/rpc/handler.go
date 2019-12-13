package rpc

import (
	"context"
	"github.com/tinyhole/im/ap/tcpserver/client"
	"github.com/tinyhole/im/idl/mua/im/ap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	tcpClient client.Client
}

func NewHandler(cli client.Client) *Handler {
	return &Handler{tcpClient: cli}
}

func (h *Handler) PushMsg(cxt context.Context, req *ap.PushMsgReq, rsp *ap.PushMsgRsp) (err error) {
	err = h.tcpClient.Push(req.Fid, req.SrvName, req.MethodName, req.Data)
	if err != nil {
		return status.Error(codes.Internal, "push msg failed")
	}

	return nil
}
