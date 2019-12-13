package entity

import "time"

const (
	PRT_Unknown  = 0
	PRT_Followed = 1 //
	PRT_Friended = 2 //
	PRT_Blocked  = 3
)

type PersonalRelation struct {
	//ID       string   `bson:"_id"`       //关系实例ID
	SrcUID   int64    `bson:"src_uid"`   //关系主UID
	DstUID   int64    `bson:"dst_uid"`   //目标uid
	IsFollow bool     `bson"is_follow"`  //关注标识
	IsFriend bool     `bson:"is_friend"` // 好友标识
	IsBlock  bool     `bson:"is_block"`  // 拉黑标识
	Remark   string   `bson:"remark"`    //备注名称
	Tag      []string `bson:"tag"`       //标签
	Desc     string   `bson:"desc"`      //描述
	FollowAt int64    `bson:"follow_at"` //关注时间
	FriendAt int64    `bson:"friend_at"`
	BlockAt  int64    `bson:"block_at"`
	CreateAt int64    `bson:"create_at"`
}

func NewPersonalRelation(srcUID, dstUID int64) *PersonalRelation {
	return &PersonalRelation{
		SrcUID:   srcUID,
		DstUID:   dstUID,
		IsFollow: false,
		IsFriend: false,
		IsBlock:  false,
		Remark:   "",
		Tag:      nil,
		Desc:     "",
		FollowAt: 0,
		FriendAt: 0,
		BlockAt:  0,
		CreateAt: time.Now().Unix(),
	}
}

func (p *PersonalRelation) Follow() {
	p.IsFollow = true
	p.FollowAt = time.Now().Unix()
}

//Unfollow 取消关注时，自动消除好友关系
func (p *PersonalRelation) Unfollow() {
	p.IsFollow = false
	p.IsFriend = false
}

func (p *PersonalRelation) Friend() {
	p.FriendAt = time.Now().Unix()
	p.IsFriend = true
}

func (p *PersonalRelation) Unfriend() {
	p.IsFriend = false
}

func (p *PersonalRelation) SetRemark(remark string) {
	p.Remark = remark
}

func (p *PersonalRelation) AppendTag(tag string) {
	p.Tag = append(p.Tag, tag)
}

func (p *PersonalRelation) Block() {
	p.BlockAt = time.Now().Unix()
	p.IsBlock = true
}

func (p *PersonalRelation) Unblock() {
	p.IsBlock = false
}
