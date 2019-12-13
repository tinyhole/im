package objconv

import (
	"github.com/tinyhole/im/idl/mua/im/logic"
	"github.com/tinyhole/im/logic/domain/valueobj"
)

type sessionConv struct{}

var (
	SessionConv = sessionConv{}
)

func (s sessionConv) DO2DTO(src *valueobj.SessionInfo) *logic.SessionInfo {
	if src == nil {
		return nil
	}

	return &logic.SessionInfo{
		Uid:        src.UID,
		Fid:        src.ApFid,
		ApName:     src.ApName,
		ApID:       src.ApID,
		DeviceType: src.DeviceTye,
	}
}

func (s sessionConv) SliceDO2DTO(src []*valueobj.SessionInfo) []*logic.SessionInfo {
	if src == nil {
		return nil
	}
	rets := make([]*logic.SessionInfo, len(src), len(src))
	for idx, itr := range src {
		rets[idx] = s.DO2DTO(itr)
	}

	return rets
}
