package nsq

import (
	"github.com/pkg/errors"
	"github.com/tinyhole/im/logic/application/eventbus"
	"github.com/tinyhole/im/logic/infrastructure/logger"
	"github.com/youzan/go-nsq"
	"time"
)

type Options struct {
	maxInFlight       int
	nsqLookupdAddr    []string
	heartbeatInterval time.Duration
}

type Option func(o *Options)

func WithMaxInFlight(maxInFight int) Option {
	return func(o *Options) {
		o.maxInFlight = maxInFight
	}
}

func WithNSQLookupdAddr(addr []string) Option {
	return func(o *Options) {
		o.nsqLookupdAddr = addr
	}
}

func WithHeartbeatInterval(interval time.Duration) Option {
	return func(o *Options) {
		o.heartbeatInterval = interval
	}
}

func NewOptions() *Options {
	return &Options{
		maxInFlight:       10000,
		nsqLookupdAddr:    []string{"127.0.0.1:4161"},
		heartbeatInterval: time.Second,
	}
}

type NSQConsumer struct {
	options  *Options
	topic    string
	channel  string
	handler  eventbus.Handler
	log      logger.Logger
	consumer *nsq.Consumer
}

func NewConsumer(topic, channel string, handler eventbus.Handler, log logger.Logger, opts ...Option) (*NSQConsumer, error) {
	options := NewOptions()
	for _, o := range opts {
		o(options)
	}
	nsqConsumer := &NSQConsumer{
		topic:   topic,
		channel: channel,
		handler: handler,
		options: options,
		log:     log,
	}
	conf := nsq.NewConfig()
	conf.HeartbeatInterval = time.Second
	conf.MaxInFlight = 1000
	consumer, err := nsq.NewConsumer(topic, channel, conf)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	consumer.AddHandler(nsqConsumer)
	consumer.SetLogger(nsqConsumer, nsq.LogLevelError)
	nsqConsumer.consumer = consumer
	return nsqConsumer, nil
}

func (c *NSQConsumer) Connect() (err error) {
	err = c.consumer.ConnectToNSQLookupds(c.options.nsqLookupdAddr)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *NSQConsumer) HandleMessage(message *nsq.Message) error {
	return c.handler.Handle(message.Body)
}

func (c *NSQConsumer) Output(calldepth int, s string) error {
	c.log.Error(s)
	return nil
}

func (c *NSQConsumer) Stop() {
	c.consumer.Stop()
	<-c.consumer.StopChan
}
