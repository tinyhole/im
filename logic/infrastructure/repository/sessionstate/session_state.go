package sessionstate

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/tinyhole/im/logic/domain/repository"
	"github.com/tinyhole/im/logic/domain/valueobj"
)

type sessionStateRepo struct {
	pool *redis.Pool
}

func NewSessionStateRepo(pool *redis.Pool) repository.SessionStateRepository {
	return &sessionStateRepo{
		pool: pool,
	}
}

func (s *sessionStateRepo) key(uid int64, apSrvID int32, apSessionID int64) string {
	return fmt.Sprintf("session|%d|%d|%d", uid, apSrvID, apSessionID)
}

func (s *sessionStateRepo) prefixKey(uid int64) string {
	return fmt.Sprintf("session|%d|*", uid)
}

func (s *sessionStateRepo) List(uid int64) ([]*valueobj.SessionInfo, error) {
	var (
		err    error
		prefix string
		reply  interface{}
		keys   []string
		datas  []string
		rets   []*valueobj.SessionInfo
	)
	con := s.pool.Get()
	defer con.Close()
	prefix = s.prefixKey(uid)
	reply, err = con.Do("keys", prefix)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	fmt.Printf("========>[%v]", reply)
	keys, err = redis.Strings(reply, err)
	args := []interface{}{}
	for _, itr := range keys {
		args = append(args, itr)
	}
	if len(args) == 0 {
		return nil, repository.ErrNotFound
	}
	reply, err = con.Do("mget", args...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	datas, err = redis.Strings(reply, err)
	if err != nil {
		return nil, err
	}

	for _, itr := range datas {
		tmp := valueobj.SessionInfo{}
		err = json.Unmarshal([]byte(itr), &tmp)
		if err != nil {
			continue
		}
		rets = append(rets, &tmp)
	}

	return rets, err
}

func (s *sessionStateRepo) Save(sessionState *valueobj.SessionInfo) error {
	var (
		err  error
		data []byte
		key  string
	)

	key = s.key(sessionState.UID, sessionState.ApID, sessionState.ApFid)
	data, err = json.Marshal(sessionState)
	if err != nil {
		return errors.WithStack(err)
	}
	con := s.pool.Get()
	defer con.Close()
	_, err = con.Do("setex", key, 30, data)
	return err
}

func (s *sessionStateRepo) Delete(uid int64, apSrvID int32, apSessionID int64) error {
	var (
		err error
		key string
	)
	key = s.key(uid, apSrvID, apSessionID)
	con := s.pool.Get()
	defer con.Close()

	_, err = con.Do("del", key)
	return err
}

func (s *sessionStateRepo) Refresh(uid int64, apSrvID int32, apSessionID int64) error {
	var (
		key string
		err error
	)

	key = s.key(uid, apSrvID, apSessionID)
	con := s.pool.Get()
	defer con.Close()
	_, err = con.Do("expire", key, 60)
	return err
}

func (s *sessionStateRepo) BatchList(uids []int64) (rets []*valueobj.SessionInfo, err error) {
	var (
		ret []*valueobj.SessionInfo
	)
	for _, itr := range uids {
		ret, err = s.List(itr)
		if err != nil {
			return
		}
		rets = append(rets, ret...)
	}

	return
}
