package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/tinyhole/im/idl/mua/im/job"
)

type NotifyObserver struct{}

func (n *NotifyObserver) Call(cli Client, data []byte) {
	notify := job.MsgNotify{}
	proto.Unmarshal(data, &notify)
	fmt.Printf("receiv a notify [%v]\n", notify)
	req := job.PullMsgReq{
		InboxID: notify.InboxID,
		Seq:     notify.Seq,
	}
	pbReq, err := proto.Marshal(&req)
	if err != nil {
		return
	}
	rsp, err := Call("mua.im.job", "Job.PullMsg", pbReq)
	if err != nil {
		fmt.Printf("cli.Call error [%v]\n", err)
	}
	msg := job.PullMsgRsp{}
	proto.Unmarshal(rsp, &msg)
	fmt.Printf("rsp=============> [%v]\n", msg)
	return
}
