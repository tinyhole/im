package repository

import (
	"github.com/tinyhole/im/relation/domain/entity"
	"github.com/tinyhole/im/relation/domain/repository"
	"gopkg.in/mgo.v2/bson"
)

type GroupRelationRepo struct {
	table string
	db    *MgoDB
}

func NewGroupRelationRepo(db *MgoDB) repository.GroupRelation {
	return &GroupRelationRepo{
		table: "group_relation",
		db:    db,
	}
}

func (g *GroupRelationRepo) Get(srcUID, groupID int64) (*entity.GroupRelation, error) {
	var (
		err error
		ret entity.GroupRelation
	)
	query := bson.M{
		"src_uid":     srcUID,
		"group_id":    groupID,
		"is_validate": true,
	}
	err = g.db.One(g.table, query, &ret)
	return &ret, err
}

func (g *GroupRelationRepo) ListMember(groupID int64, page, pageSize int32) ([]*entity.GroupRelation, int, error) {
	var (
		err   error
		rets  []*entity.GroupRelation
		total int
	)
	query := bson.M{
		"group_id":    groupID,
		"is_validate": true,
	}

	total, err = g.db.List(g.table, query, []string{"-role", "-join_at"}, page, pageSize, &rets)
	return rets, total, err

}

func (g *GroupRelationRepo) ListGroup(uid int64, page, pageSize int32) ([]*entity.GroupRelation, int, error) {
	var (
		err   error
		rets  []*entity.GroupRelation
		total int
	)
	query := bson.M{
		"src_uid":     uid,
		"is_validate": true,
	}
	total, err = g.db.List(g.table, query, []string{"-role", "-join_at"}, page, pageSize, &rets)
	return rets, total, err
}

func (g *GroupRelationRepo) Save(relation *entity.GroupRelation) error {
	var (
		err error
	)
	query := bson.M{
		"src_uid":  relation.SrcUID,
		"group_id": relation.GroupID,
	}
	update := bson.M{
		"$set": bson.M{
			"join_at":     relation.JoinAt,
			"is_validate": relation.IsValidate,
			"role":        relation.Role,
		},
		"$setOnInsert": bson.M{
			"create_at": relation.CreateAt,
		},
	}

	_, err = g.db.Upsert(g.table, query, update)
	return err
}

func (g *GroupRelationRepo) Quit(uid, groupID int64) error {
	query := bson.M{
		"src_uid":     uid,
		"group_id":    groupID,
		"is_validate": true,
	}

	update := bson.M{
		"$set": bson.M{
			"is_validate": false,
		},
	}
	return g.db.Update(g.table, query, update)
}
