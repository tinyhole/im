package nsq

import (
	"github.com/pkg/errors"
	"github.com/tinyhole/im/job/application/eventbus"
	"github.com/tinyhole/im/job/infrastructure/config"
	"github.com/tinyhole/im/job/infrastructure/logger"
	"github.com/youzan/go-nsq"
	"time"
)

type Manager struct {
	nsqLookdAddr []string
	consumers    map[string]*NSQConsumer
	producers    map[string]*NSQProducer
	nsqdAddr     string
	nsqProducer  *nsq.Producer
	log          logger.Logger
}

func NewManager(conf *config.Config, log logger.Logger) eventbus.Manager {
	return &Manager{
		nsqLookdAddr: conf.NSQLookupdAddr,
		consumers:    make(map[string]*NSQConsumer),
		producers:    make(map[string]*NSQProducer),
		nsqdAddr:     conf.NSQDAddr,
		log:          log,
	}
}

func (m *Manager) AddConsumer(handler eventbus.Handler) error {
	consumer, err := NewConsumer(handler.Topic(), handler.Chan(), handler, m.log, WithNSQLookupdAddr(m.nsqLookdAddr))
	if err != nil {
		return errors.WithStack(err)
	}
	m.consumers[handler.Topic()] = consumer
	return nil
}

func (m *Manager) NewPublisher(topic string) (eventbus.Publisher, error) {
	if m.nsqProducer == nil {
		conf := nsq.NewConfig()
		conf.HeartbeatInterval = time.Second
		p, err := nsq.NewProducer(m.nsqdAddr, conf)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		m.nsqProducer = p
	}

	producer, err := NewProducer(topic, m.nsqProducer, m.log)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	m.producers[topic] = producer
	return producer, nil
}

func (m *Manager) Run() (err error) {
	for _, v := range m.consumers {
		err = v.Connect()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) Stop() {
	for _, v := range m.consumers {
		v.Stop()
	}
	m.nsqProducer.Stop()
}
