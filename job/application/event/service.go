package event

import (
	"github.com/tinyhole/im/job/infrastructure/logger"
)

type EventService struct {
	mgr             ConsumerMgr
	msgEventHandler *MsgHandler
	log             logger.Logger
}

func NewEventService(msgHandler *MsgHandler, mgr ConsumerMgr, log logger.Logger) *EventService {
	return &EventService{
		mgr:             mgr,
		msgEventHandler: msgHandler,
		log:             log,
	}
}

func (e *EventService) Run() error {
	err := e.mgr.AddConsumer(e.msgEventHandler.Topic(), e.msgEventHandler)
	e.log.Errorf("%v", err)
	return err
}
