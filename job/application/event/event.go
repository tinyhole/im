package event

type ConsumerMgr interface {
	AddConsumer(string, Handler) error
}

type Handler interface {
	HandleMsg([]byte) error
	Topic() string
}
