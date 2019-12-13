package rpc

import (
	"context"
	"github.com/tinyhole/im/idl/mua/im/sequence"
	"github.com/tinyhole/im/sequence/application"
)

type Handler struct {
	app *application.AppService
}

func NewHandler(app *application.AppService) *Handler {
	return &Handler{
		app: app,
	}
}

func (h *Handler) GetAutoIncrID(ctx context.Context, req *sequence.GetAutoIncrIDReq, rsp *sequence.GetAutoIncrIDRsp) error {
	id, err := h.app.GetNextAutoIncrID(req.Key)
	if err != nil {
		return err
	}
	rsp.Id = id
	return nil
}
