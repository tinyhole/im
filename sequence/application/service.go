package application

import (
	"github.com/tinyhole/im/sequence/domain/autoincrement/repository"
	"github.com/tinyhole/im/sequence/infrastructure/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppService struct {
	autoIncrRepo repository.AutoIncrRepository
	logger       logger.Logger
}

func NewAppService(autoIncrRepo repository.AutoIncrRepository,
	logger logger.Logger) *AppService {
	return &AppService{
		autoIncrRepo: autoIncrRepo,
		logger:       logger,
	}
}

func (a *AppService) GetNextAutoIncrID(key string) (id int64, err error) {
	id, err = a.autoIncrRepo.NextID(key)
	if err != nil {
		a.logger.Errorf("GetNextAutoIncrID failed [%v]", err)
		return 0, status.Error(codes.Internal, err.Error())
	}

	return id, nil
}
