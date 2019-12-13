package service

import (
	"github.com/tinyhole/im/job/domain/entity"
	"github.com/tinyhole/im/job/domain/repository"
	"github.com/tinyhole/im/job/domain/service"
	"github.com/tinyhole/im/job/infrastructure/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppService struct {
	jobSvc  *service.JobService
	msgRepo repository.Inbox
	log     logger.Logger
}

func NewAppService(svc *service.JobService,
	log logger.Logger,
	msgRepo repository.Inbox) *AppService {
	return &AppService{
		jobSvc:  svc,
		log:     log,
		msgRepo: msgRepo,
	}
}

func (a *AppService) PullMsg(inboxID int64, seq int64) (msg *entity.Message, err error) {
	msg, err = a.msgRepo.Get(inboxID, seq)
	if err != nil {
		a.log.Errorf("Get failed [%v]", err)
		return nil, status.Error(codes.Internal, "pull msg failed")
	}
	return msg, err
}

func (a *AppService) SyncPrivateInboxMsg(uid int64, seq int64) (msg []*entity.Message, err error) {

	return nil, nil
}

func (a *AppService) SyncPublicInboxMsg(uid int64) (msg []*entity.Message, err error) {
	return nil, nil
}