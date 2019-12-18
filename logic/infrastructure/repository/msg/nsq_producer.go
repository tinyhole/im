package msg

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/tinyhole/im/idl/mua/im"
	"github.com/tinyhole/im/logic/domain/repository"
	"github.com/tinyhole/im/logic/infrastructure/logger"
	"github.com/youzan/go-nsq"
	"time"
)

type MsgRepo struct {
	topic    string
	producer *nsq.Producer
	log      logger.Logger
}

func NewMsgRepo(producer *nsq.Producer, log logger.Logger) repository.MsgRepository {
	ret := &MsgRepo{
		topic:    "mua.im.chat_msg",
		producer: producer,
		log:      log,
	}
	producer.SetLogger(ret, nsq.LogLevelError)
	return ret
}

func (m *MsgRepo) PushMsg(msg *im.Msg) error {

}

func (m *MsgRepo) Output(calldepth int, s string) error {
	m.log.Error(s)
	return nil
}
