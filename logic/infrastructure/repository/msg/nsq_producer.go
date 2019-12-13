package msg

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/tinyhole/im/idl/mua/im"
	"github.com/tinyhole/im/logic/domain/repository"
	"github.com/youzan/go-nsq"
	"time"
)

type MsgRepo struct {
	topic    string
	producer *nsq.Producer
}

func NewMsgRepo(producer *nsq.Producer) repository.MsgRepository {
	return &MsgRepo{
		topic:    "mua.im.chat_msg",
		producer: producer,
	}
}

func (m *MsgRepo) PushMsg(msg *im.Msg) error {
	msgData, err := proto.Marshal(msg)
	if err != nil {
		return errors.WithStack(err)
	}
	doneChan := make(chan *nsq.ProducerTransaction)
	m.producer.PublishAsync(m.topic, msgData, doneChan)

	select {
	case <-doneChan:
		return nil
	case <-time.After(5 * time.Second):
		return errors.New("send msg to nsq timeout")
	}
	return nil
}
