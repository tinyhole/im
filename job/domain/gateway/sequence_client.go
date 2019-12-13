package gateway

type SequenceClient interface {
	GetPrivateSeq(inboxID int64) (seq int64, err error)
	GetGroupSeq(inboxID int64) (seq int64, err error)
}
