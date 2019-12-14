package gateway

type SequenceClient interface {
	GetPrivateSeq(inboxID string) (seq int64, err error)
	GetGroupSeq(inboxID string) (seq int64, err error)
}
