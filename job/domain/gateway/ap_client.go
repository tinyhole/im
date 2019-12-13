package gateway

import (
	"github.com/tinyhole/im/job/domain/valueobj"
)

type ApClient interface {
	PushMsg(apID int32, fid int64, msg *valueobj.MsgNotify) error
}
