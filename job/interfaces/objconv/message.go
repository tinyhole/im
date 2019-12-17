package objconv

import (
	"github.com/tinyhole/im/idl/mua/im"
	"github.com/tinyhole/im/job/domain/entity"
)

type messageConv struct{}

var (
	MessageConv = messageConv{}
)

func (m messageConv) DTO2DO(src *im.Msg) *entity.Message {
	if src == nil {
		return nil
	}

	return &entity.Message{
		MsgID:       src.MsgID,
		SrcID:       src.SrcID,
		DstID:       src.DstID,
		MsgType:     int32(src.MsgType),
		Content:     src.Content,
		ContentType: int32(src.ContentType),
	}
}

func (m messageConv) DO2DTO(src *entity.Message) *im.Msg {
	if src == nil {
		return nil
	}

	return &im.Msg{
		MsgID:       src.MsgID,
		SrcID:       src.SrcID,
		DstID:       src.DstID,
		ContentType: im.ContentType(src.ContentType),
		Seq:         src.Seq,
		Time:        src.Time,
		Content:     src.Content,
		MsgType:     im.MsgType(src.MsgType),
		InboxID:     src.InboxID,
	}
}

func (m messageConv) SliceDO2DTO(src []*entity.Message) []*im.Msg {
	if src == nil {
		return nil
	}

	rets := make([]*im.Msg, len(src), len(src))
	for idx, itr := range src {
		rets[idx] = m.DO2DTO(itr)
	}
	return rets
}
