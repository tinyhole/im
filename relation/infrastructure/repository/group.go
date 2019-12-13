package repository

import (
	"fmt"
	"github.com/tinyhole/im/relation/domain/entity"
	"github.com/tinyhole/im/relation/domain/repository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type GroupRepo struct {
	table string
	db    *MgoDB
}

func NewGroupRepo(db *MgoDB) repository.Group {
	return &GroupRepo{
		table: "group",
		db:    db,
	}
}

func (g *GroupRepo) GetSeqName() string {
	return fmt.Sprintf("%s.%s", g.db.DBName, g.table)
}

func (g *GroupRepo) Save(data *entity.Group) error {
	var (
		err error
	)
	query := bson.M{
		"_id": data.ID,
	}
	update := bson.M{
		"$set": bson.M{
			"name":   data.Name,
			"notice": data.Notice,
			"type":   data.Type,
		},
		"$setOnInsert": bson.M{
			"create_at": data.CreatedAt,
		},
	}
	_, err = g.db.Upsert(g.table, query, update)
	return err
}

func (g *GroupRepo) Get(id int64) (*entity.Group, error) {
	var (
		err error
		ret entity.Group
	)
	query := bson.M{
		"_id": id,
	}
	err = g.db.One(g.table, query, &ret)
	if err == mgo.ErrNotFound {
		return nil, repository.ErrNotFound
	}
	return &ret, err
}

func (g *GroupRepo) Search(name string, page, pageSize int32) ([]*entity.Group, int, error) {
	var (
		err   error
		rets  []*entity.Group
		total int
	)

	ses := g.db.Session.Copy()
	if ses == nil {
		return nil, 0, ErrMgoSessIsNil
	}

	query := bson.M{
		"$regex": "*name*",
	}

	total, err = ses.DB(g.db.DBName).C(g.table).Find(query).Count()
	if err != nil {
		return nil, 0, err
	}
	offset, limit := calcPage(int(page), int(pageSize))
	err = ses.DB(g.db.DBName).C(g.table).Find(query).Skip(offset).Limit(limit).All(&rets)
	return rets, total, err
}
