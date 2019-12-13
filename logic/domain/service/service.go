package service

import (
	"github.com/pkg/errors"
	gateway2 "github.com/tinyhole/im/logic/domain/gateway"
	valueobj2 "github.com/tinyhole/im/logic/domain/valueobj"
	"time"
)

var (
	ErrAuthFailed = errors.New("auth failed")
)

type SessionService struct {
	authCli gateway2.AuthClient
}

func NewSessionService(authCli gateway2.AuthClient) *SessionService {

	return &SessionService{authCli: authCli}
}

func (s *SessionService) Auth(uid int64, token string, apSrvID int32, apFid int64,
	remoteIP string, deviceType int32) (state *valueobj2.SessionInfo, err error) {
	var (
		authFlag bool
	)

	authFlag, err = s.authCli.Auth(uid, token)

	if authFlag == false {
		return nil, ErrAuthFailed
	}

	state = &valueobj2.SessionInfo{
		UID:       uid,
		DeviceTye: deviceType,
		RemoteIP:  remoteIP,
		ApID:      apSrvID,
		ApFid:     apFid,
		UpdateAt:  time.Now().Unix(),
	}

	return state, nil
}
