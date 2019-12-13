package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/tinyhole/im/idl/mua/im"
	"github.com/tinyhole/im/idl/mua/im/logic"
	"os"
	"os/signal"
	"strconv"
)

var (
	cli Client
)

func main() {
	cli = NewClient("")
	tmpUid, _ := strconv.Atoi(os.Args[1])
	uid := int64(tmpUid)
	cli.Init(int64(uid), "")

	msg := im.Msg{
		SrcID:       uid,
		DstID:       2,
		ContentType: im.ContentType_ContentTypeText,
		Content:     "hello world",
		MsgType:     im.MsgType_MsgTypePrivte,
	}

	req := logic.PushMsgReq{
		Msg: &msg,
	}
	pbReq, _ := proto.Marshal(&req)
	_, err := cli.Call("mua.im.logic", "Logic.PushMsg", pbReq)
	if err != nil {
		fmt.Printf("call Logic.PushMsg failed [%v]", err.Error())
	}

	ob := NotifyObserver{}
	cli.AddOB("mua.im.job", "Job.PushMsg", &ob)
	wait()
}

func wait() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
