package repository

import "github.com/tinyhole/im/idl/mua/im"

type MsgRepository interface {
	PushMsg(msg *im.Msg) error
}
