package service

import (
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/tinyhole/im/idl/mua/im"
	eventbus2 "github.com/tinyhole/im/logic/application/eventbus"
	"github.com/tinyhole/im/logic/domain/repository"
	"github.com/tinyhole/im/logic/domain/service"
	"github.com/tinyhole/im/logic/domain/valueobj"
	"github.com/tinyhole/im/logic/infrastructure/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppService struct {
	sessionStateRepo repository.SessionStateRepository
	sessionSvc       *service.SessionService
	eventBusMgr      eventbus2.Manager
	logger           logger.Logger
	publisher        eventbus2.Publisher
}

func NewAppService(stateRepository repository.SessionStateRepository,
	sessionSvc *service.SessionService,
	eventBusMgr eventbus2.Manager,
	logger logger.Logger) (*AppService, error) {
	ret := &AppService{
		sessionStateRepo: stateRepository,
		sessionSvc:       sessionSvc,
		logger:           logger,
		eventBusMgr:      eventBusMgr,
	}
	publisher, err := ret.eventBusMgr.NewPublisher("mua.im.chat_msg")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ret.publisher = publisher
	return ret, nil
}

func (a *AppService) Ping(uid int64, apSrvID int32, apSessionID int64) (err error) {
	err = a.sessionStateRepo.Refresh(uid, apSrvID, apSessionID)
	if err != nil {
		a.logger.Errorf("sessionStateRepo.Refresh failed [%v]", err)
		err = status.Error(codes.Internal, "ping error")
	}

	return nil
}

func (a *AppService) SignIn(uid int64, deviceType int32, token string,
	apSrvID int32, apSessionID int64, remoteIP string) error {
	var (
		sessionState *valueobj.SessionInfo
		err          error
	)
	sessionState, err = a.sessionSvc.Auth(uid, token, apSrvID, apSessionID, remoteIP, deviceType)
	if err != nil {

		return err
	}
	err = a.sessionStateRepo.Save(sessionState)
	if err != nil {
		a.logger.Errorf("save session state error [%v]", err)
		return status.Error(codes.Internal, "save session failed")
	}
	return nil
}

func (a *AppService) PushMsg(msg *im.Msg) (err error) {
	byteMsg, err := proto.Marshal(msg)
	if err != nil {
		a.logger.Errorf("marshal msg failed [%v]", err)
		return status.Error(codes.InvalidArgument, "marshal msg failed")
	}
	err = a.publisher.AsyncPublish(byteMsg)
	if err != nil {
		a.logger.Errorf("push msg error [%v]", err)
		return status.Error(codes.Internal, "send msg error")
	}
	return nil
}

func (a *AppService) ListSessionInfo(uid int64) ([]*valueobj.SessionInfo, error) {
	rets, err := a.sessionStateRepo.List(uid)
	if err != nil {
		a.logger.Errorf("sessionStateRepo.List failed [%v]", err)
		err = status.Error(codes.Internal, "list session info failed")
	}

	return rets, err
}

func (a *AppService) BatchListSessionInfo(uids []int64) ([]*valueobj.SessionInfo, error) {
	rets, err := a.sessionStateRepo.BatchList(uids)
	if err != nil {
		a.logger.Errorf("sessionStateRepo.BatchList failed [%v]", err)
		err = status.Error(codes.Internal, "batch list session info failed")
	}

	return rets, err
}
