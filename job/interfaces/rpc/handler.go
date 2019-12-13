package rpc

import (
	"context"
	"github.com/tinyhole/im/idl/mua/im/job"
	"github.com/tinyhole/im/job/application/service"
	"github.com/tinyhole/im/job/interfaces/objconv"
)

type Handler struct {
	appSvc *service.AppService
}

func NewHandler(app *service.AppService) *Handler {
	return &Handler{appSvc: app}
}

func (h *Handler) PullMsg(ctx context.Context, req *job.PullMsgReq, rsp *job.PullMsgRsp) error {
	msg, err := h.appSvc.PullMsg(req.InboxID, req.Seq)
	if err != nil {
		return err
	}
	rsp.Msg = objconv.MessageConv.DO2DTO(msg)

	return nil
}

func (h *Handler) SyncPrivateInboxMsg(ctx context.Context, req *job.SyncPrivateInboxMsgReq, rsp *job.SyncPrivateInboxMsgRsp) error {
	return nil
}

func (h *Handler) SyncPublicInboxMsg(ctx context.Context, req *job.SyncPublicInboxMsgReq, rsp *job.SyncPublicInboxMsgRsp) error {
	return nil
}
