package service

import "errors"

var (
	ErrUnknownMsgType       = errors.New("unknown message type")
	ErrProcessMsgFailed     = errors.New("process message failed")
	ErrGenerateNotifyFailed = errors.New("generate notify failed")
)
