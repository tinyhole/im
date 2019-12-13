package nsq

import (
	"github.com/tinyhole/im/logic/infrastructure/config"
	"github.com/youzan/go-nsq"
	"time"
)

func NewProducer(conf *config.Config) (producer *nsq.Producer, err error) {
	nsqConf := nsq.NewConfig()
	nsqConf.HeartbeatInterval = time.Second
	producer, err = nsq.NewProducer(conf.NSQDAddr, nsqConf)
	return
}
