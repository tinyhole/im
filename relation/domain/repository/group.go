package repository

import (
	"github.com/tinyhole/im/relation/domain/entity"
)

type Group interface {
	Get(int64) (*entity.Group, error)
	Search(name string, page, pageSize int32) ([]*entity.Group, int, error)
	Save(data *entity.Group) error
	GetSeqName() string
}
