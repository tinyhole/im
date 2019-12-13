package repository

import (
	"github.com/tinyhole/im/logic/domain/valueobj"
)

type SessionStateRepository interface {
	List(uid int64) ([]*valueobj.SessionInfo, error)
	BatchList(uid []int64) ([]*valueobj.SessionInfo, error)
	Save(sessionState *valueobj.SessionInfo) error
	Delete(uid int64, apSrvID int32, apSessionID int64) error
	Refresh(uid int64, apSrvID int32, apSessionID int64) error
}
