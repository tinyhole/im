package gateway

import (
	"github.com/tinyhole/im/job/domain/valueobj"
)

type SessionClient interface {
	ListSessionInfo(uid int64) (info []*valueobj.SessionInfo, err error)
	BatchListSessionInfo([]int64) ([]*valueobj.SessionInfo, error)
}
