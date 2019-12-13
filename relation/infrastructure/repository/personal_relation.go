package repository

import (
	"github.com/tinyhole/im/relation/domain/entity"
	"github.com/tinyhole/im/relation/domain/repository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PersonalRelationRepo struct {
	table string
	db    *MgoDB
}

func NewPersonalRelationRepo(db *MgoDB) repository.PersonalRelation {
	return &PersonalRelationRepo{
		table: "personal_relation",
		db:    db,
	}
}

func (p *PersonalRelationRepo) Save(data *entity.PersonalRelation) error {
	var (
		err error
	)
	query := bson.M{
		"src_uid": data.SrcUID,
		"dst_uid": data.DstUID,
	}
	update := bson.M{
		"$set": bson.M{
			"is_follow": data.IsFollow,
			"is_friend": data.IsFriend,
			"is_block":  data.IsBlock,
			"remark":    data.Remark,
			"desc":      data.Desc,
			"tag":       data.Tag,
			"follow_at": data.FollowAt,
			"friend_at": data.FriendAt,
			"block_at":  data.BlockAt,
		},
		"$setOnInsert": bson.M{
			"create_at": data.CreateAt,
		},
	}

	_, err = p.db.Upsert(p.table, query, update)
	return err
}

func (p *PersonalRelationRepo) Get(srcUID int64, dstUID int64) (*entity.PersonalRelation, error) {
	var (
		ret entity.PersonalRelation
		err error
	)
	query := bson.M{
		"src_uid": srcUID,
		"dst_uid": dstUID,
	}

	err = p.db.One(p.table, query, &ret)
	return &ret, err
}

func (p *PersonalRelationRepo) ListFriend(srcUID int64, page, pageSize int32) ([]*entity.PersonalRelation, int, error) {
	var (
		total int
		rets  []*entity.PersonalRelation
		err   error
	)
	query := bson.M{
		"src_uid":   srcUID,
		"is_friend": true,
	}
	total, err = p.db.List(p.table, query, []string{"-friend_at"}, page, pageSize, &rets)
	return rets, total, err
}

func (p *PersonalRelationRepo) ListFollow(srcUID int64, page, pageSize int32) ([]*entity.PersonalRelation, int, error) {
	var (
		total int
		rets  []*entity.PersonalRelation
		err   error
	)
	query := bson.M{
		"src_uid":   srcUID,
		"is_follow": true,
	}
	total, err = p.db.List(p.table, query, []string{"-follow_at"}, page, pageSize, &rets)
	return rets, total, err
}

func (p *PersonalRelationRepo) ListBlock(srcUID int64, page, pageSize int32) ([]*entity.PersonalRelation, int, error) {
	var (
		total int
		rets  []*entity.PersonalRelation
		err   error
	)

	query := bson.M{
		"src_uid":  srcUID,
		"is_block": true,
	}

	total, err = p.db.List(p.table, query, []string{"-block_at"}, page, pageSize, &rets)
	return rets, total, err
}

func (p *PersonalRelationRepo) GetFriend(srcUID, dstUID int64) (*entity.PersonalRelation, error) {
	var (
		err error
		ret entity.PersonalRelation
	)
	query := bson.M{
		"src_uid":   srcUID,
		"dst_uid":   dstUID,
		"is_friend": true,
	}
	err = p.db.One(p.table, query, &ret)
	if err == mgo.ErrNotFound {
		return nil, repository.ErrNotFound
	}
	return &ret, err
}

func (p *PersonalRelationRepo) GetFollow(srcUID, dstUID int64) (*entity.PersonalRelation, error) {
	var (
		err error
		ret entity.PersonalRelation
	)
	query := bson.M{
		"src_uid":   srcUID,
		"dst_uid":   dstUID,
		"is_follow": true,
	}
	err = p.db.One(p.table, query, &ret)
	if err == mgo.ErrNotFound {
		return nil, repository.ErrNotFound
	}
	return &ret, err
}

func (p *PersonalRelationRepo) GetBlock(srcUID, dstUID int64) (*entity.PersonalRelation, error) {
	var (
		err error
		ret entity.PersonalRelation
	)

	query := bson.M{
		"src_uid":  srcUID,
		"dst_uid":  dstUID,
		"is_block": true,
	}
	err = p.db.One(p.table, query, &ret)
	if err == mgo.ErrNotFound {
		return nil, repository.ErrNotFound
	}

	return &ret, nil
}

func (p *PersonalRelationRepo) Unfriend(srcUID, dstUID int64) error {
	query := bson.M{
		"src_uid":   srcUID,
		"dst_uid":   dstUID,
		"is_friend": true,
	}

	update := bson.M{
		"$set": bson.M{
			"is_friend": false,
		},
	}

	return p.db.Update(p.table, query, update)
}

func (p *PersonalRelationRepo) Unfollow(srcUID, dstUID int64) error {
	query := bson.M{
		"src_uid":   srcUID,
		"dst_uid":   dstUID,
		"is_follow": true,
	}

	update := bson.M{
		"$set": bson.M{
			"is_follow": false,
		},
	}

	return p.db.Update(p.table, query, update)
}

func (p *PersonalRelationRepo) Unblock(srcUID, dstUID int64) error {
	query := bson.M{
		"src_uid":  srcUID,
		"dst_uid":  dstUID,
		"is_block": true,
	}
	update := bson.M{
		"$set": bson.M{
			"is_block": false,
		},
	}
	return p.db.Update(p.table, query, update)
}
