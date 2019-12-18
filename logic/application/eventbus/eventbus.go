package eventbus

type Handler interface {
	Handle([]byte) error
	Topic() string
	Chan() string
}

type Publisher interface {
	AsyncPublish([]byte) error
}

type Manager interface {
	AddConsumer(handler Handler) error
	NewPublisher(topic string) (Publisher, error)
	Run() error
	Stop()
}
