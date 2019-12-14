package util

import "fmt"

type InboxIDClient interface {
	GetPersonalInboxID(int64)string
	GetGroupInboxID(int64)string
}


type inboxIDCli struct{}

func NewInboxIDClient()InboxIDClient{
	return &inboxIDCli{}
}

func(i *inboxIDCli)GetPersonalInboxID(uid int64)string{
	return fmt.Sprintf("%d@p",uid)
}

func (i *inboxIDCli)GetGroupInboxID(gid int64)string{
	return fmt.Sprintf("%d@g",gid)
}
