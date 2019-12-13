package repository

import (
	"errors"
	"gopkg.in/mgo.v2"
)

var (
	ErrMgoSessIsNil = errors.New("mongo session is nil")
)

type MgoDB struct {
	DBName  string
	Session *mgo.Session
}

func NewMgoDB(session *mgo.Session) *MgoDB {
	return &MgoDB{
		DBName:  "mua_im_relation",
		Session: session,
	}
}

func (m *MgoDB) Insert(c string, docs ...interface{}) error {
	ses := m.Session.Copy()
	if ses == nil {
		return ErrMgoSessIsNil
	}
	defer ses.Close()
	return ses.DB(m.DBName).C(c).Insert(docs...)
}

func (m *MgoDB) Update(c string, query interface{}, update interface{}) error {
	ses := m.Session.Copy()
	if ses == nil {
		return ErrMgoSessIsNil
	}
	defer ses.Close()
	return ses.DB(m.DBName).C(c).Update(query, update)
}

func (m *MgoDB) Upsert(c string, query interface{}, update interface{}) (*mgo.ChangeInfo, error) {
	ses := m.Session.Copy()
	if ses == nil {
		return nil, ErrMgoSessIsNil
	}
	defer ses.Close()
	return ses.DB(m.DBName).C(c).Upsert(query, update)
}

func (m *MgoDB) One(c string, query interface{}, ret interface{}) error {
	ses := m.Session.Copy()
	if ses == nil {
		return ErrMgoSessIsNil
	}
	defer ses.Close()
	return ses.DB(m.DBName).C(c).Find(query).One(ret)
}

func (m *MgoDB) All(c string, query interface{}, rets interface{}) error {
	ses := m.Session.Copy()
	if ses == nil {
		return ErrMgoSessIsNil
	}
	defer ses.Close()
	return ses.DB(m.DBName).C(c).Find(query).All(rets)
}

func (m *MgoDB) List(c string, query interface{}, sortFields []string, page, pageSize int32, rets interface{}) (int, error) {
	ses := m.Session.Copy()
	if ses == nil {
		return 0, ErrMgoSessIsNil
	}
	defer ses.Close()
	cnt, err := ses.DB(m.DBName).C(c).Find(query).Count()
	if err != nil {
		return 0, err
	}
	offset, limit := calcPage(int(page), int(pageSize))
	err = ses.DB(m.DBName).C(c).Find(query).Sort(sortFields...).Skip(offset).Limit(limit).All(rets)
	return cnt, err
}

func (m *MgoDB) Remove(c string, query interface{}) error {
	ses := m.Session.Copy()
	if ses == nil {
		return ErrMgoSessIsNil
	}
	defer ses.Close()
	return ses.DB(m.DBName).C(c).Remove(query)
}
