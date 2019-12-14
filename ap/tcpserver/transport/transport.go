package transport

import (
	"github.com/tinyhole/im/ap/tcpserver/protocol/pack"
)

type Transport interface {
	//Dial(addr string, opts ...DialOption)(Client, error)
	Listen(addr string, opts ...ListenOption) (Listener, error)
}

type Socket interface {
	Recv() *pack.ApPackage
	Send(apPackage *pack.ApPackage)error
	Close() error
	Local() string
	Remote() string
	ID() uint32
	GetAuthState() bool
	UpdateAuthState(state bool)
	SetUID(int64)
	GetUID() int64
}

type Listener interface {
	Addr() string
	Close() error
	Accept(setUp, destroy, heartbeat func(socket Socket)) error
}
