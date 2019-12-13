package repository

import (
	"github.com/tinyhole/im/relation/domain/entity"
)

type GroupRelation interface {
	ListMember(groupID int64, page, pageSize int32) ([]*entity.GroupRelation, int, error)
	ListGroup(uid int64, page, pageSize int32) ([]*entity.GroupRelation, int, error)
	Save(relation *entity.GroupRelation) error
	Quit(uid int64, groupID int64) error
	Get(uid int64, groupID int64) (*entity.GroupRelation, error)
}
