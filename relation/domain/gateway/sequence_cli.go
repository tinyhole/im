package gateway

type SequenceClient interface {
	GenerateGroupID(key string) (id int64, err error)
}
