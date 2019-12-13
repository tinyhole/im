// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: mua/im_sequence.proto

/*
Package sequence is a generated protocol buffer package.

It is generated from these files:
	mua/im_sequence.proto

It has these top-level messages:
	GetAutoIncrIDReq
	GetAutoIncrIDRsp
*/
package sequence

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

// Client API for Sequence service

type SequenceService interface {
	GetAutoIncrID(ctx context.Context, in *GetAutoIncrIDReq, opts ...client.CallOption) (*GetAutoIncrIDRsp, error)
}

type sequenceService struct {
	c    client.Client
	name string
}

func NewSequenceService(name string, c client.Client) SequenceService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "mua.im.sequence"
	}
	return &sequenceService{
		c:    c,
		name: name,
	}
}

func (c *sequenceService) GetAutoIncrID(ctx context.Context, in *GetAutoIncrIDReq, opts ...client.CallOption) (*GetAutoIncrIDRsp, error) {
	req := c.c.NewRequest(c.name, "Sequence.GetAutoIncrID", in)
	out := new(GetAutoIncrIDRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sequence service

type SequenceHandler interface {
	GetAutoIncrID(context.Context, *GetAutoIncrIDReq, *GetAutoIncrIDRsp) error
}

func RegisterSequenceHandler(s server.Server, hdlr SequenceHandler, opts ...server.HandlerOption) error {
	type sequence interface {
		GetAutoIncrID(ctx context.Context, in *GetAutoIncrIDReq, out *GetAutoIncrIDRsp) error
	}
	type Sequence struct {
		sequence
	}
	h := &sequenceHandler{hdlr}
	return s.Handle(s.NewHandler(&Sequence{h}, opts...))
}

type sequenceHandler struct {
	SequenceHandler
}

func (h *sequenceHandler) GetAutoIncrID(ctx context.Context, in *GetAutoIncrIDReq, out *GetAutoIncrIDRsp) error {
	return h.SequenceHandler.GetAutoIncrID(ctx, in, out)
}