package util

import (
	"github.com/google/uuid"
	"gopkg.in/mgo.v2/bson"
)

type MsgIDClient interface {
	GetMsgID() string
}

type UUIDCli struct{}

func NewUUIDCli() MsgIDClient {
	return &UUIDCli{}
}

func (u *UUIDCli) GetMsgID() string {
	return uuid.New().String()
}

type ObjectIDCli struct{}

func NewObjectIDCli() MsgIDClient {
	return &ObjectIDCli{}
}

func (o *ObjectIDCli) GetMsgID() string {
	return bson.NewObjectId().Hex()
}
