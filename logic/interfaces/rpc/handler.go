package rpc

import (
	"context"
	"github.com/tinyhole/im/idl/mua/im/logic"
	"github.com/tinyhole/im/logic/application/service"
	"github.com/tinyhole/im/logic/interfaces/objconv"
)

type Handler struct {
	appSvc *service.AppService
}

func NewHandler(appSvc *service.AppService) *Handler {
	return &Handler{appSvc: appSvc}
}

func (h *Handler) PushMsg(ctx context.Context, req *logic.PushMsgReq, rsp *logic.PushMsgRsp) (err error) {
	return h.appSvc.PushMsg(req.Msg)
}

func (h *Handler) Ping(ctx context.Context, req *logic.PingReq, rsp *logic.PingRsp) (err error) {
	return h.appSvc.Ping(req.Uid, req.ApID, req.ApFid)
}

func (h *Handler) SignIn(ctx context.Context, req *logic.SignInReq, rsp *logic.SignInRsp) (err error) {
	return h.appSvc.SignIn(req.Uid, req.DeviceType, req.Token, req.ApID, req.ApFid, req.RemoteIP)
}

func (h *Handler) ListSessionInfo(ctx context.Context, req *logic.ListSessionInfoReq, rsp *logic.ListSessionInfoRsp) (err error) {
	rets, err := h.appSvc.ListSessionInfo(req.Uid)
	if err != nil {
		return err
	}
	rsp.Infos = objconv.SessionConv.SliceDO2DTO(rets)
	return
}

func (h *Handler) BatchListSessionInfo(ctx context.Context, req *logic.BatchListSessionInfoReq,
	rsp *logic.BatchListSessionInfoRsp) (err error) {
	rets, err := h.appSvc.BatchListSessionInfo(req.Uids)
	if err != nil {
		return err
	}
	rsp.Infos = objconv.SessionConv.SliceDO2DTO(rets)
	return

}
