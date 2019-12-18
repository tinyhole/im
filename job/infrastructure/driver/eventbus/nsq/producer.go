package nsq

import (
	"github.com/pkg/errors"
	"github.com/tinyhole/im/job/infrastructure/logger"
	"github.com/youzan/go-nsq"
	"time"
)

type NSQProducer struct {
	nsqProducer *nsq.Producer
	topic       string
	log         logger.Logger
}

func NewProducer(topic string, p *nsq.Producer, log logger.Logger) (*NSQProducer, error) {
	ret := &NSQProducer{
		nsqProducer: p,
		topic:       topic,
		log:         log,
	}
	p.SetLogger(ret, nsq.LogLevelError)
	return ret, nil
}

func (p *NSQProducer) Output(calldepth int, s string) error {
	p.log.Error(s)
	return nil
}

func (p *NSQProducer) AsyncPublish(data []byte) error {
	doneChan := make(chan *nsq.ProducerTransaction)
	p.nsqProducer.PublishAsync(p.topic, data, doneChan)

	select {
	case <-doneChan:
		return nil
	case <-time.After(5 * time.Second):
		return errors.New("send msg to nsq timeout")
	}
	return nil
}

func (p *NSQProducer) Stop() {
	p.nsqProducer.Stop()
}
