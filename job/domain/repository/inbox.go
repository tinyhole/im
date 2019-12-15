package repository

import (
	"github.com/tinyhole/im/job/domain/entity"
)

type Inbox interface {
	Save(data *entity.Message) error
	Get(inboxID string, seq int64) (msg *entity.Message, err error)
	List(inboxID string, seq, startTime, endTime int64, page, pageSize int32) (rets []*entity.Message, total int, err error)
}
