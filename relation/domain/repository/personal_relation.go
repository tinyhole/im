package repository

import (
	"github.com/tinyhole/im/relation/domain/entity"
)

type PersonalRelation interface {
	Save(*entity.PersonalRelation) error
	Get(srcUID int64, dstUID int64) (*entity.PersonalRelation, error)
	ListFriend(srcUID int64, page, pageSize int32) ([]*entity.PersonalRelation, int, error)
	ListFollow(srcUID int64, page, pageSize int32) ([]*entity.PersonalRelation, int, error)
	ListBlock(srcUID int64, page, pageSize int32) ([]*entity.PersonalRelation, int, error)
	GetFriend(srcUID, dstUID int64) (*entity.PersonalRelation, error)
	GetFollow(srcUID, dstUID int64) (*entity.PersonalRelation, error)
	GetBlock(srcUID, dstUID int64) (*entity.PersonalRelation, error)
	//Unfriend(srcUID, dstUID int64) error
	//Unfollow(srcUID, dstUID int64) error
	//Unblock(srcUID, dstUID int64) error
}
