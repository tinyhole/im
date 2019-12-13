package entity

import (
	"github.com/pkg/errors"
	"time"
)

const (
	GR_GeneralMember = 1 //普通成员
	GR_Owner         = 2 //所有者
	GR_Manager       = 3 //管理员
)

var (
	ErrInvalidateRole = errors.New("invalidate relation")
)

type GroupRelation struct {
	//ID      string `bson:"_id"`      //关系实例ID
	SrcUID     int64 `bson:"src_uid"`     //
	GroupID    int64 `bson:"group_id"`    //
	JoinAt     int64 `bson:"join_at"`     //加入时间
	Role       int32 `bson:"role"`        //角色
	IsValidate bool  `bson:"is_validate"` //是否有效
	CreateAt   int64 `bson:"create_at"`
}

func NewGroupRelation(srcUID, groupID int64) *GroupRelation {
	return &GroupRelation{
		SrcUID:     srcUID,
		GroupID:    groupID,
		JoinAt:     0,
		Role:       GR_GeneralMember,
		IsValidate: false,
		CreateAt:   time.Now().Unix(),
	}
}

func (g *GroupRelation) Join() {
	g.JoinAt = time.Now().Unix()
	g.IsValidate = true
}

func (g *GroupRelation) ChangeRole(role int32) error {
	switch role {
	case GR_GeneralMember:
	case GR_Owner:
	case GR_Manager:
		g.Role = role
	default:
		return ErrInvalidateRole
	}

	return nil
}

func (g *GroupRelation) Quit() {
	g.Role = GR_GeneralMember
	g.IsValidate = false
}
