// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: mua/im_ap.proto

/*
Package ap is a generated protocol buffer package.

It is generated from these files:
	mua/im_ap.proto

It has these top-level messages:
	PushMsgReq
	PushMsgRsp
*/
package ap

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AP service

type APService interface {
	PushMsg(ctx context.Context, in *PushMsgReq, opts ...client.CallOption) (*PushMsgRsp, error)
}

type aPService struct {
	c    client.Client
	name string
}

func NewAPService(name string, c client.Client) APService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "mua.im.ap"
	}
	return &aPService{
		c:    c,
		name: name,
	}
}

func (c *aPService) PushMsg(ctx context.Context, in *PushMsgReq, opts ...client.CallOption) (*PushMsgRsp, error) {
	req := c.c.NewRequest(c.name, "AP.PushMsg", in)
	out := new(PushMsgRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AP service

type APHandler interface {
	PushMsg(context.Context, *PushMsgReq, *PushMsgRsp) error
}

func RegisterAPHandler(s server.Server, hdlr APHandler, opts ...server.HandlerOption) error {
	type aP interface {
		PushMsg(ctx context.Context, in *PushMsgReq, out *PushMsgRsp) error
	}
	type AP struct {
		aP
	}
	h := &aPHandler{hdlr}
	return s.Handle(s.NewHandler(&AP{h}, opts...))
}

type aPHandler struct {
	APHandler
}

func (h *aPHandler) PushMsg(ctx context.Context, in *PushMsgReq, out *PushMsgRsp) error {
	return h.APHandler.PushMsg(ctx, in, out)
}
