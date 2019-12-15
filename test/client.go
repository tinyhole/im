package main

import (
	"container/list"
	"fmt"
	"github.com/dubbogo/getty"
	"github.com/pkg/errors"
	"github.com/tinyhole/im/ap/tcpserver/protocol/pack"
	"github.com/tinyhole/im/ap/tcpserver/protocol/tcprpc"
	"go.uber.org/atomic"
	"net"
	"sync"
	"time"
)

const (
	CronPeriod = 6e9
)

type Client interface {
	Call(srvName string, methodName string, req []byte) (rsp []byte, err error)
	//Notify(srvName string,methodName string, req []byte)
	Init(uid int64, token string)
	AddOB(srvName, endPoint string, observer Observer)
}

type Observer interface {
	Call(cli Client, data []byte)
}

type Response struct {
	ch  chan (struct{})
	rsp *pack.ApPackage
}

type client struct {
	con       getty.Client
	socket    *tcpTransportSocket
	msgPool   map[int64]*Response
	poolMutex sync.Mutex
	seq       atomic.Int64
	obPool    map[string]*list.List
	uid       int64
	token     string
}

func NewClient(addr string) Client {
	return &client{
		msgPool: make(map[int64]*Response),
		obPool:  make(map[string]*list.List),
	}
}

func (c *client) Call(srvName string, endpoint string, req []byte) ([]byte, error) {
	var (
		rsp []byte
	)
	seq := c.seq.Add(1)
	c.poolMutex.Lock()
	response := &Response{
		ch:  make(chan struct{}),
		rsp: nil,
	}
	c.msgPool[seq] = response
	c.poolMutex.Unlock()
	reqPack := &pack.ApPackage{
		Header: &pack.Header{
			Request: &pack.RequestMeta{
				ServiceName: srvName,
				Endpoint:    endpoint,
				CallType:    pack.CallType_Sync,
			},
			Auth: &pack.AuthInfo{
				Uid:   c.uid,
				Token: c.token,
			},
			Seq: seq,
			Device: &pack.Device{
				Guid: "111",
				Type: 1,
			},
		},
		Body: req,
	}
	c.socket.Write(reqPack)
	<-response.ch
	if response.rsp.Header.Response.ErrCode != 0 {
		err := &RspError{
			ErrCode: response.rsp.Header.Response.ErrCode,
			ErrText: response.rsp.Header.Response.ErrText,
		}
		return nil, err
	}

	c.poolMutex.Lock()
	if c.msgPool[seq].rsp.Body != nil {
		rsp = c.msgPool[seq].rsp.Body
	}
	delete(c.msgPool, seq)
	c.poolMutex.Unlock()
	return rsp, nil
}

func (c *client) Notify(srvName string, endPoint string, req []byte) {
	topic := fmt.Sprintf("%s.%s", srvName, endPoint)
	if v, ok := c.obPool[topic]; ok {
		for ele := v.Front(); ele != nil; ele = ele.Next() {
			ele.Value.(Observer).Call(c, req)
		}
	}
}

func (c *client) AddOB(srvName, endPoint string, ob Observer) {
	topic := fmt.Sprintf("%s.%s", srvName, endPoint)
	if v, ok := c.obPool[topic]; ok {
		v.PushBack(ob)
	} else {
		c.obPool[topic] = list.New()
		c.obPool[topic].PushBack(ob)
	}
}

func (c *client) Init(uid int64, token string) {
	c.uid = uid
	c.token = token
	c.con = getty.NewTCPClient(getty.WithServerAddress("127.0.0.1:8080"), getty.WithConnectionNumber(1))
	socket := &tcpTransportSocket{session: nil, client: c, uid: uid, token: token}
	c.con.RunEventLoop(func(session getty.Session) error {
		var (
			ok      bool
			tcpConn *net.TCPConn
		)
		if tcpConn, ok = session.Conn().(*net.TCPConn); !ok {
			panic(fmt.Sprintf("%s, session.conn{%#v} is not tcp connection\n", session.Stat(), session.Conn()))
		}

		tcpConn.SetNoDelay(true)
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(time.Duration(time.Second * 6))
		tcpConn.SetReadBuffer(262144)
		tcpConn.SetWriteBuffer(65536)

		session.SetName(fmt.Sprintf("tcp-%s", session.RemoteAddr()))
		session.SetMaxMsgLen(65536)
		session.SetRQLen(1024)
		session.SetWQLen(1024)
		session.SetReadTimeout(time.Second)
		session.SetWriteTimeout(time.Second * 5)
		session.SetCronPeriod(int(CronPeriod) / 1e6) //6 second
		session.SetWaitTime(time.Second * 7)
		//session.SetTaskPool(t.taskPool)

		session.SetPkgHandler(&tcpTransport{})
		session.SetEventListener(socket)
		return nil
	})
	c.socket = socket

}

type tcpTransport struct{}

func (t *tcpTransport) Read(session getty.Session, data []byte) (interface{}, int, error) {
	apPack := &pack.ApPackage{}
	unCodec := tcprpc.GetCodec()
	defer tcprpc.PutCodec(unCodec)
	err := unCodec.Unmarshal(data, apPack)
	if err != nil {
		if err == tcprpc.ErrTotalLengthNotEnough || err == tcprpc.ErrFlagLengthNotEnough {
			return nil, 0, nil
		}
		return nil, 0, errors.WithStack(err)
	}
	totalLen := int(tcprpc.HeadLenBytesSize + tcprpc.TotalLenBytesSize + unCodec.(*tcprpc.TcpCodec).TotalLen)
	return apPack, totalLen, err
}

func (t *tcpTransport) Write(session getty.Session, pkg interface{}) ([]byte, error) {
	unCodec := tcprpc.GetCodec()
	defer tcprpc.PutCodec(unCodec)
	data, err := unCodec.Marshal(pkg)
	return data, err
}

type tcpTransportSocket struct {
	session getty.Session
	client  *client
	uid     int64
	token   string
}

func (t *tcpTransportSocket) OnOpen(session getty.Session) error {
	t.session = session
	return nil
}

func (t *tcpTransportSocket) OnError(session getty.Session, err error) {
	t.session.Close()
}

//OnClose 回调上层错误处理函数，关闭msgchan
func (t *tcpTransportSocket) OnClose(session getty.Session) {
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
		if pbPkg.Header.Request.ServiceName == "mua.im.ap" &&
			pbPkg.Header.Request.Endpoint == "AP.ping" {
			fmt.Println("ping")
			return
		}

		if pbPkg.Header.Request.ServiceName == "mua.im.ap" &&
			pbPkg.Header.Request.Endpoint == "AP.pong" {
			fmt.Println("pong")
			return
		}

		go t.client.Notify(pbPkg.Header.Request.ServiceName,
			pbPkg.Header.Request.Endpoint, pbPkg.Body)

	}

	if pbPkg.Header.Response != nil {
		fmt.Printf("rsp [ %v\n", pbPkg.Header)
		t.client.poolMutex.Lock()
		if response, ok := t.client.msgPool[pbPkg.Header.Seq]; ok {
			response.rsp = pbPkg
			response.ch <- struct{}{}
		}
		t.client.poolMutex.Unlock()
	}
}

func (t *tcpTransportSocket) OnCron(session getty.Session) {
	var (
		err error
	)
	req := &pack.ApPackage{
		Header: &pack.Header{
			Request: &pack.RequestMeta{
				ServiceName: "mua.im.ap",
				Endpoint:    "AP.Ping",
				CallType:    pack.CallType_Push,
			},
			Auth: &pack.AuthInfo{
				Uid:   t.uid,
				Token: t.token,
			},
			Device: &pack.Device{
				Guid: "111",
				Type: 1,
			},
			Seq: 0,
		},
	}
	err = session.WritePkg(req, time.Duration(5*time.Second))

	if err != nil {
		session.Close()
		return
	}

	//active := session.GetActive()
	/*if CronPeriod < time.Since(active).Nanoseconds() {
		session.Close()
	}
	*/
}

func (t *tcpTransportSocket) Write(data *pack.ApPackage) error {
	err := t.session.WritePkg(data, time.Second*5)
	return err
}
