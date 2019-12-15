package repository

import (
	"github.com/pkg/errors"
	"github.com/tinyhole/im/job/domain/entity"
	"github.com/tinyhole/im/job/domain/repository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type inboxRepo struct {
	table string
	db    *MgoDB
}

func NewInboxRepo(db *MgoDB) repository.Inbox {
	return &inboxRepo{
		table: "inbox",
		db:    db,
	}
}

func (i *inboxRepo) Get(inboxID string, seq int64) (*entity.Message, error) {
	var (
		err error
		ret entity.Message
	)
	query := bson.M{
		"inbox_id": inboxID,
		"seq":      seq,
	}
	err = i.db.One(i.table, query, &ret)
	if err == mgo.ErrNotFound {
		err = repository.ErrNotFound
	}
	if err != nil && err != mgo.ErrNotFound {
		err = errors.WithStack(err)
	}
	return &ret, err
}

func (i *inboxRepo) Save(data *entity.Message) (err error) {
	query := bson.M{
		"msg_id":   data.MsgID,
		"inbox_id": data.InboxID,
	}

	update := bson.M{
		"$setOnInsert": bson.M{
			"src_id":       data.SrcID,
			"dst_id":       data.DstID,
			"msg_type":     data.MsgType,
			"seq":          data.Seq,
			"time":         data.Time,
			"content":      data.Content,
			"content_type": data.ContentType,
		},
	}

	_, err = i.db.Upsert(i.table, query, update)

	return err
}

func (i *inboxRepo) List(inboxID string, seq, startTime, endTime int64, page, pageSize int32) (rets []*entity.Message, total int, err error) {
	query := bson.M{
		"inbox_id": inboxID,
		"time": bson.M{
			"$gte": startTime,
			"$lte": endTime,
		},
		"seq": bson.M{
			"$gt": seq,
		},
	}

	total, err = i.db.List(i.table, query, []string{"+seq"}, page, pageSize, &rets)
	return
}
