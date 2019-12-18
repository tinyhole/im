package event

import (
	eventbus2 "github.com/tinyhole/im/job/application/eventbus"
	"github.com/tinyhole/im/job/infrastructure/logger"
)

type EventService struct {
	mgr             eventbus2.Manager
	msgEventHandler *MsgHandler
	log             logger.Logger
}

func NewEventService(msgHandler *MsgHandler, mgr eventbus2.Manager, log logger.Logger) *EventService {
	return &EventService{
		mgr:             mgr,
		msgEventHandler: msgHandler,
		log:             log,
	}
}

func (e *EventService) Run() error {
	err := e.mgr.AddConsumer(e.msgEventHandler)
	e.log.Errorf("add consumer error [%v]", err)
	err = e.mgr.Run()
	return err
}

func (e *EventService) Stop() {
	e.mgr.Stop()
}
