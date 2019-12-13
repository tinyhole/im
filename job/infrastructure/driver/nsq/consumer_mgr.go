package nsq

import (
	"github.com/pkg/errors"
	"github.com/tinyhole/im/job/application/event"
	"github.com/tinyhole/im/job/infrastructure/config"
	"github.com/tinyhole/im/job/infrastructure/logger"
)

type ConsumerMgr struct {
	nsqLookAddr string
	channelName string
	consumers   map[string]*NSQConsumer
	nsqdAddr    string
	log         logger.Logger
}

func NewConsumerMgr(conf *config.Config, log logger.Logger) event.ConsumerMgr {
	return &ConsumerMgr{
		nsqLookAddr: conf.NSQLookUpAddr,
		channelName: conf.NSQChannel,
		consumers:   make(map[string]*NSQConsumer),
		nsqdAddr:    conf.NSQDAddr,
		log:         log,
	}
}

func (c *ConsumerMgr) AddConsumer(topic string, handler event.Handler) error {
	consumer, err := NewConsumer(topic, c.channelName, handler, c.log, WithNSQLookupdAddr(c.nsqLookAddr), WithNSQDAddr(c.nsqdAddr))
	if err != nil {
		return errors.WithStack(err)
	}
	c.consumers[topic] = consumer
	return nil
}
