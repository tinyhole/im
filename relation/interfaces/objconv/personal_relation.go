package objconv

import (
	"github.com/tinyhole/im/idl/mua/im/relation"
	"github.com/tinyhole/im/relation/domain/entity"
)

type PersonalRelationConv struct{}

func NewPersonalRelationConv() *PersonalRelationConv {
	return &PersonalRelationConv{}
}

func (p *PersonalRelationConv) DO2DTO(src *entity.PersonalRelation) *relation.PersonalRelation {
	if src == nil {
		return nil
	}
	ret := &relation.PersonalRelation{
		SrcUID:   src.SrcUID,
		DstUID:   src.DstUID,
		IsFollow: src.IsFollow,
		IsFriend: src.IsFriend,
		IsBlock:  src.IsBlock,
		Remark:   src.Remark,
		Tag:      src.Tag,
		Desc:     src.Desc,
		FollowAt: src.FollowAt,
		FriendAt: src.FriendAt,
		BlockAt:  src.BlockAt,
	}
	return ret
}

func (p *PersonalRelationConv) DTO2DO(src *relation.PersonalRelation) *entity.PersonalRelation {
	if src == nil {
		return nil
	}

	ret := &entity.PersonalRelation{
		SrcUID:   src.SrcUID,
		DstUID:   src.DstUID,
		IsFollow: src.IsFollow,
		IsFriend: src.IsFriend,
		IsBlock:  src.IsBlock,
		Remark:   src.Remark,
		Tag:      src.Tag,
		Desc:     src.Desc,
		FollowAt: src.FollowAt,
		FriendAt: src.FriendAt,
		BlockAt:  src.BlockAt,
	}

	return ret
}

func (p *PersonalRelationConv) SliceDO2DTO(src []*entity.PersonalRelation) []*relation.PersonalRelation {
	rets := make([]*relation.PersonalRelation, len(src), len(src))
	for idx, itr := range src {
		rets[idx] = p.DO2DTO(itr)
	}
	return rets
}
