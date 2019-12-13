package objconv

import (
	"github.com/tinyhole/im/idl/mua/im/relation"
	"github.com/tinyhole/im/relation/domain/entity"
)

type GroupConv struct{}

func NewGroupConv() *GroupConv {
	return &GroupConv{}
}

func (g *GroupConv) DO2DTO(src *entity.Group) *relation.Group {
	if src == nil {
		return nil
	}
	ret := &relation.Group{
		ID:        src.ID,
		Name:      src.Name,
		Notice:    src.Notice,
		CreatedAt: src.CreatedAt,
		Type:      relation.GroupType(src.Type),
	}

	return ret
}

func (g *GroupConv) DTO2DO(src *relation.Group) *entity.Group {
	if src == nil {
		return nil
	}

	ret := &entity.Group{
		ID:        src.ID,
		Name:      src.Name,
		Notice:    src.Notice,
		CreatedAt: src.CreatedAt,
		Type:      int32(src.Type),
	}

	return ret
}

func (g *GroupConv) SliceDO2DTO(src []*entity.Group) []*relation.Group {
	rets := make([]*relation.Group, len(src), len(src))
	for idx, itr := range src {
		rets[idx] = g.DO2DTO(itr)
	}

	return rets
}
