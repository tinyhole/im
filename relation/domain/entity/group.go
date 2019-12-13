package entity

import "time"

const (
	GT_Temporary = 1 //临时
	GT_General   = 2 //普通
	GT_Super     = 3 //超级
)

type Group struct {
	ID        int64  `bson:"_id"`    //群组ID
	Name      string `bson:"name"`   //群名称
	Notice    string `bson:"notice"` //群公告
	CreatedAt int64  `bson:"create_at"`
	Type      int32  `bson:"type"`
}

func NewGroup(name string) *Group {
	return &Group{
		ID:        0,
		Name:      name,
		Notice:    "",
		Type:      GT_General,
		CreatedAt: time.Now().Unix(),
	}
}
