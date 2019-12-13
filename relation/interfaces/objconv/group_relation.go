package objconv

import (
	"github.com/tinyhole/im/idl/mua/im/relation"
	"github.com/tinyhole/im/relation/domain/entity"
)

type GroupRelationConv struct{}

func NewGroupRelationConv() *GroupRelationConv {
	return &GroupRelationConv{}
}

func (g *GroupRelationConv) DO2DTO(src *entity.GroupRelation) *relation.GroupRelation {
	if src == nil {
		return nil
	}

	ret := &relation.GroupRelation{
		SrcUID:  src.SrcUID,
		GroupID: src.GroupID,
		JoinAt:  src.JoinAt,
		Role:    relation.GroupRoleType(src.Role),
	}

	return ret
}

func (g *GroupRelationConv) DTO2DO(src *relation.GroupRelation) *entity.GroupRelation {
	if src == nil {
		return nil
	}
	ret := &entity.GroupRelation{
		SrcUID:  src.SrcUID,
		GroupID: src.GroupID,
		JoinAt:  src.JoinAt,
		Role:    int32(src.Role),
	}

	return ret
}

func (g *GroupRelationConv) SliceDO2DTO(src []*entity.GroupRelation) []*relation.GroupRelation {
	rets := make([]*relation.GroupRelation, len(src), len(src))
	for idx, itr := range src {
		rets[idx] = g.DO2DTO(itr)
	}

	return rets
}
