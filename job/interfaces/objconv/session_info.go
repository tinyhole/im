package objconv

import (
	"github.com/tinyhole/im/idl/mua/im/logic"
	"github.com/tinyhole/im/job/domain/valueobj"
)

type sessionInfoConv struct{}

var (
	SessionInfoConv = sessionInfoConv{}
)

func (s sessionInfoConv) DTO2DO(src *logic.SessionInfo) *valueobj.SessionInfo {
	if src == nil {
		return nil
	}
	return &valueobj.SessionInfo{
		Uid:        src.Uid,
		Fid:        src.Fid,
		ApID:       src.ApID,
	}
}

func (s sessionInfoConv) SliceDTO2DO(src []*logic.SessionInfo) []*valueobj.SessionInfo {
	if src == nil {
		return nil
	}

	rets := make([]*valueobj.SessionInfo, len(src), len(src))
	for idx, itr := range src {
		rets[idx] = s.DTO2DO(itr)
	}
	return rets
}
