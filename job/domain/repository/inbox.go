package repository

import (
	"github.com/tinyhole/im/job/domain/entity"
)

type Inbox interface {
	Save(data *entity.Message) error
	Get(inboxID int64, seq int64) (msg *entity.Message, err error)
}
