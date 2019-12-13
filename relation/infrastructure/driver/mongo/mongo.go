package mongo

import (
	"errors"
	"github.com/tinyhole/im/relation/infrastructure/config"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	ErrMgoAddrIncorrect = errors.New("mongo address incorrect")
)

func NewMgoSession(cfg *config.Config) (*mgo.Session, error) {
	if len(cfg.MgoAddrs) == 0 {
		return nil, ErrMgoAddrIncorrect
	}
	dialInfo := &mgo.DialInfo{
		Addrs:          cfg.MgoAddrs,
		Direct:         false,
		Timeout:        time.Second * 3,
		FailFast:       false,
		ReplicaSetName: cfg.MgoReplicaSetName,
		Username:       cfg.MgoUser,
		Password:       cfg.MgoPassword,
		PoolLimit:      cfg.MgoPoolLimit,
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return session, nil
}
