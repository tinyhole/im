package gateway

type SequenceCli interface {
	GetSeq(key string) (id int64, err error)
}
