package transport

import (
	"github.com/dubbogo/getty"
	"github.com/tinyhole/im/ap/tcpserver/protocol/pack"
	"time"
)

const (
	CronPeriod = 6e9
)

type tcpTransportSocket struct {
	session     getty.Session
	msgChan     chan *pack.ApPackage
	authState   bool
	uid         int64
	destroyFn   func(socket Socket)
	heartbeatFn func(socket Socket)
}

func NewTcpTransportSocket(session getty.Session, destroyFn, heartbeat func(socket Socket)) getty.EventListener {
	return &tcpTransportSocket{
		session:     session,
		msgChan:     make(chan *pack.ApPackage, 1024),
		authState:   false,
		destroyFn:   destroyFn,
		heartbeatFn: heartbeat,
	}
}

func (t *tcpTransportSocket) GetAuthState() bool {
	return t.authState
}

func (t *tcpTransportSocket) SetUID(uid int64) {
	t.uid = uid
}

func (t *tcpTransportSocket) GetUID() int64 {
	return t.uid
}

func (t *tcpTransportSocket) UpdateAuthState(state bool) {
	t.authState = state
}

func (t *tcpTransportSocket) Recv() *pack.ApPackage {
	pkg, ok := <-t.msgChan
	if ok {
		return pkg
	}
	return nil
}

func (t *tcpTransportSocket) Send(intrepidPackage *pack.ApPackage) {
	var (
		err error
	)
	err = t.session.WritePkg(intrepidPackage, time.Second*5)
	if err != nil {
		t.session.Close()
		return
	}
}

func (t *tcpTransportSocket) Close() error {
	t.session.Close()

	return nil
}

func (t *tcpTransportSocket) Local() string {

	return t.session.LocalAddr()
}

func (t *tcpTransportSocket) Remote() string {
	return t.session.RemoteAddr()
}

func (t *tcpTransportSocket) ID() uint32 {
	return t.session.ID()
}

func (t *tcpTransportSocket) OnOpen(session getty.Session) error {
	//t.session = session
	return nil
}

func (t *tcpTransportSocket) OnError(session getty.Session, err error) {
	t.session.Close()
}

//OnClose 回调上层错误处理函数，关闭msgchan
func (t *tcpTransportSocket) OnClose(session getty.Session) {
	if t.destroyFn != nil {
		t.destroyFn(t)
	}
	close(t.msgChan)
	t.session.Close()
}

func (t *tcpTransportSocket) OnMessage(session getty.Session, pkg interface{}) {
	var (
		pbPkg *pack.ApPackage
		ok    bool
	)

	pbPkg, ok = pkg.(*pack.ApPackage)
	if !ok {
		return
	}

	if pbPkg.Header.Request != nil {
		if pbPkg.Header.Request.ServiceName == "ap" &&
			pbPkg.Header.Request.Endpoint == "ping" {
			if t.heartbeatFn != nil {
				t.heartbeatFn(t)
			}
			return
		}
	}

	t.msgChan <- pbPkg
}

func (t *tcpTransportSocket) OnCron(session getty.Session) {
	var (
		err error
	)
	req := &pack.ApPackage{
		Header: &pack.Header{
			Request: &pack.RequestMeta{
				ServiceName: "ap",
				Endpoint:    "pong",
				CallType:    pack.CallType_Push,
			},
			Seq: 0,
		},
	}
	err = session.WritePkg(req, time.Duration(5*time.Second))

	if err != nil {
		session.Close()
		return
	}

	/*
		active := session.GetActive()
		if CronPeriod < time.Since(active).Nanoseconds() {
			session.Close()
		}
	*/
}
