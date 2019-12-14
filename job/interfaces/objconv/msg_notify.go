package objconv

import (
	"github.com/tinyhole/im/idl/mua/im/job"
	"github.com/tinyhole/im/job/domain/valueobj"
)

type msgNotifyConv struct{}

var (
	MsgNotifyConv = msgNotifyConv{}
)

func (m msgNotifyConv) DO2DTO(src *valueobj.Notify) *job.MsgNotify {
	if src == nil {
		return nil
	}
	return &job.MsgNotify{
		InboxID: src.InboxID,
		Seq:     src.Seq,
	}
}
