package entity

const (
	MsgypeUnknown = 0
	MsgTypePrivte = 1 //单聊
	MsgTypeGroup  = 2 //群聊
)

type Message struct {
	MsgID       string `bson:"msg_id"`
	InboxID     string `bson:"inbox_id"` //收件箱id
	SrcID       int64  `bson:"src_id"`
	DstID       int64  `bson:"dst_id"`
	MsgType     int32  `bson:"msg_type"` //消息类型，私聊，群聊
	Seq         int64  `bson:"seq"`
	Time        int64  `bson:"time"`
	Content     string `bson:"content"`
	ContentType int32  `bson:"content_type"` //内容类型
}

func (m *Message) CopyToInbox(inboxID string) *Message {
	return &Message{
		MsgID:       m.MsgID,
		InboxID:     inboxID,
		SrcID:       m.SrcID,
		DstID:       m.DstID,
		MsgType:     m.MsgType,
		Seq:         0,
		Time:        m.Time,
		Content:     m.Content,
		ContentType: m.ContentType,
	}
}
