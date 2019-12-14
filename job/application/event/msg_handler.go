package event

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/tinyhole/im/idl/mua/im"
	"github.com/tinyhole/im/job/domain/gateway"
	"github.com/tinyhole/im/job/domain/repository"
	"github.com/tinyhole/im/job/domain/service"
	"github.com/tinyhole/im/job/domain/valueobj"
	"github.com/tinyhole/im/job/infrastructure/logger"
	"github.com/tinyhole/im/job/interfaces/objconv"
)

type MsgHandler struct {
	svc *service.JobService
	log logger.Logger
	repo repository.Inbox
	notifyClient  gateway.ApClient
}

func NewMsgHandler(svc *service.JobService,
	log logger.Logger,
	repo repository.Inbox,
	notifyClient gateway.ApClient) *MsgHandler {
	return &MsgHandler{
		svc: svc,
		log: log,
		repo:repo,
		notifyClient:notifyClient,
	}
}

func (m *MsgHandler) HandleMsg(data []byte) (err error) {
	rawMsg := im.Msg{}
	err = proto.Unmarshal(data, &rawMsg)
	msg := objconv.MessageConv.DTO2DO(&rawMsg)
	rets, notifies,err := m.svc.ProcessMsg(msg)
	if err != nil{
		if errors.Cause(err) == service.ErrProcessMsgFailed{
			return err
		}
		return nil
	}
	for _, itr := range rets{
		err = m.repo.Save(itr)
	}
	for _, itr := range notifies {
		if itr == nil{
			fmt.Println("=======")
		}
		pbNotify := objconv.MsgNotifyConv.DO2DTO(itr.Notify)
		data ,_:= proto.Marshal(pbNotify)
		for k,v := range itr.SessMap {
			if v.Len() > 1{
				tmpFids := []int64{}
				for item := v.Front();item != nil; item.Next(){
					tmpFids = append(tmpFids, item.Value.(*valueobj.SessionInfo).Fid)
				}
				m.notifyClient.Broadcast(k, tmpFids, data)
			}else if v.Len() == 1{
				m.notifyClient.Unicast(k,v.Front().Value.(*valueobj.SessionInfo).Fid, data)
			}
		}
	}
	return
}

func (m *MsgHandler) Topic() string {
	return "mua.im.chat_msg"
}
