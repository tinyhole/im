package event

import (
	"github.com/golang/protobuf/proto"
	"github.com/tinyhole/im/idl/mua/im"
	"github.com/tinyhole/im/job/domain/service"
	"github.com/tinyhole/im/job/infrastructure/logger"
	"github.com/tinyhole/im/job/interfaces/objconv"
)

type MsgHandler struct {
	svc *service.JobService
	log logger.Logger
}

func NewMsgHandler(svc *service.JobService,
	log logger.Logger) *MsgHandler {
	return &MsgHandler{
		svc: svc,
		log: log,
	}
}

func (m *MsgHandler) HandleMsg(data []byte) (err error) {
	rawMsg := im.Msg{}
	err = proto.Unmarshal(data, &rawMsg)
	msg := objconv.MessageConv.DTO2DO(&rawMsg)
	m.log.Debugf("receive msg [%v]", msg)
	err = m.svc.ProcessMsg(msg)
	if err != nil {
		m.log.Errorf("%v", err)
	}
	return
}

func (m *MsgHandler) Topic() string {
	return "mua.im.chat_msg"
}
