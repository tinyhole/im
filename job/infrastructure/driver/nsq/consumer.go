package nsq

import (
	"github.com/youzan/go-nsq"
	"github.com/pkg/errors"
	"github.com/tinyhole/im/job/application/event"
	"github.com/tinyhole/im/job/infrastructure/logger"
	"time"
)

type Options struct {
	maxInFlight    int
	nsqLookupdAddr string
	nsqDAddr       string

	heartbeatInterval time.Duration
}

type Option func(o *Options)

func WithMaxInFlight(maxInFight int) Option {
	return func(o *Options) {
		o.maxInFlight = maxInFight
	}
}

func WithNSQLookupdAddr(addr string) Option {
	return func(o *Options) {
		o.nsqLookupdAddr = addr
	}
}

func WithHeartbeatInterval(interval time.Duration) Option {
	return func(o *Options) {
		o.heartbeatInterval = interval
	}
}

func WithNSQDAddr(addr string) Option {
	return func(o *Options) {
		o.nsqDAddr = addr
	}
}

func NewOptions() *Options {
	return &Options{
		maxInFlight:       10000,
		nsqLookupdAddr:    "127.0.0.1:4161",
		nsqDAddr:          "127.0.0.1:4150",
		heartbeatInterval: time.Second,
	}
}

type NSQConsumer struct {
	options *Options
	topic   string
	channel string
	handler event.Handler
	log     logger.Logger
}

func NewConsumer(topic, channel string, handler event.Handler, log logger.Logger, opts ...Option) (*NSQConsumer, error) {
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
	err = consumer.ConnectToNSQLookupd(options.nsqLookupdAddr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return nsqConsumer, nil
}

func (c *NSQConsumer) HandleMessage(message *nsq.Message) error {
	return c.handler.HandleMsg(message.Body)
}

func (c *NSQConsumer) Output(calldepth int, s string) error {
	c.log.Error(s)
	return nil
}
