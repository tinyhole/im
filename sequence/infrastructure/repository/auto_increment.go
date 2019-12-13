package repository

import (
	"github.com/tinyhole/im/sequence/domain/autoincrement/entity"
	"github.com/tinyhole/im/sequence/domain/autoincrement/repository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type autoIncrRepo struct {
	table string
	db    *MgoDB
}

func NewAutoIncrRepo(db *MgoDB) repository.AutoIncrRepository {
	return &autoIncrRepo{
		table: "auto_increment",
		db:    db,
	}
}

func (a *autoIncrRepo) NextID(idx string) (int64, error) {

	var (
		ret entity.AutoIncrement
		err error
	)
	ses := a.db.Session.Copy()
	if ses == nil {
		return 0, ErrMgoSessIsNil
	}
	defer ses.Close()
	query := bson.M{
		"_id": idx,
	}
	update := bson.M{
		"$inc": bson.M{
			"current_id": 1,
		},
	}

	change := mgo.Change{
		Update:    update,
		Upsert:    true,
		Remove:    false,
		ReturnNew: true,
	}

	_, err = ses.DB(a.db.DBName).C(a.table).Find(query).Apply(change, &ret)

	return ret.CurrentID, err
}
