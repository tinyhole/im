package application

import (
	"github.com/pkg/errors"
	"github.com/tinyhole/im/relation/domain/entity"
	"github.com/tinyhole/im/relation/domain/logger"
	"github.com/tinyhole/im/relation/domain/repository"
	"github.com/tinyhole/im/relation/domain/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInternal = status.Error(codes.Internal, "internal error")
)

type AppService struct {
	personalRepo      repository.PersonalRelation
	groupRelationRepo repository.GroupRelation
	groupRepo         repository.Group
	relationSvc       service.RelationService
	log               logger.Logger
}

func NewAppService(personalRepo repository.PersonalRelation,
	groupRelationRepo repository.GroupRelation,
	groupRepo repository.Group,
	relationSvc service.RelationService,
	log logger.Logger) *AppService {
	return &AppService{
		personalRepo:      personalRepo,
		groupRelationRepo: groupRelationRepo,
		groupRepo:         groupRepo,
		relationSvc:       relationSvc,
		log:               log,
	}
}

func (a *AppService) GetPersonalRelation(srcUID, dstUID int64) (pr *entity.PersonalRelation, err error) {
	pr, err = a.personalRepo.Get(srcUID, dstUID)
	if err != nil {
		a.log.Errorf("%v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return
}

func (a *AppService) ListPersonalRelation(srcUID int64, typ, page, pageSize int32) ([]*entity.PersonalRelation, int, error) {
	var (
		err       error
		total     int
		relations []*entity.PersonalRelation
	)
	switch typ {
	case entity.PRT_Blocked:
		relations, total, err = a.personalRepo.ListBlock(srcUID, page, pageSize)
	case entity.PRT_Followed:
		relations, total, err = a.personalRepo.ListFollow(srcUID, page, pageSize)
	case entity.PRT_Friended:
		relations, total, err = a.personalRepo.ListFriend(srcUID, page, pageSize)
	default:
		return relations, total, status.Error(codes.InvalidArgument, "unknown relation type")
	}
	if err != nil {
		a.log.Warnf("%v", err)
		err = status.Error(codes.Internal, err.Error())
	}
	return relations, total, err
}

//Follow  todo:事物保证
func (a *AppService) Follow(srcUID, dstUID int64) error {
	var (
		srcRelation *entity.PersonalRelation
		dstRelation *entity.PersonalRelation
		err         error
	)
	srcRelation, err = a.personalRepo.GetFollow(srcUID, dstUID)
	if err != nil && err != repository.ErrNotFound {
		a.log.Errorf("personalRep.GetFollow failed %v", err)
		return ErrInternal
	}

	if err != nil && err == repository.ErrNotFound {
		srcRelation = entity.NewPersonalRelation(srcUID, dstUID)
	}

	dstRelation, err = a.personalRepo.GetFollow(dstUID, srcUID)
	if err != nil && err != repository.ErrNotFound {
		a.log.Errorf("personalRepo.GetFollow failed %v", err)
		return ErrInternal
	}

	a.relationSvc.Follow(srcRelation, dstRelation)

	err = a.personalRepo.Save(srcRelation)
	if err != nil {
		a.log.Errorf("personalRepo.Save failed %v", err)
		return ErrInternal
	}
	err = a.personalRepo.Save(dstRelation)
	if err != nil {
		a.log.Errorf("personalRepo.Save failed %v", err)
		return ErrInternal
	}
	return nil
}

//Unfollow todo:事物保证
func (a *AppService) Unfollow(srcUID, dstUID int64) error {
	var (
		srcRelation *entity.PersonalRelation
		dstRelation *entity.PersonalRelation
		err         error
	)
	srcRelation, err = a.personalRepo.GetFollow(srcUID, dstUID)
	//没有对应关系的情况下直接返回
	if err != nil && err == repository.ErrNotFound {
		a.log.Warnf("repo.GetFollow not found")
		return status.Error(codes.NotFound, "not found")
	}

	if err != nil {
		return ErrInternal
	}

	dstRelation, err = a.personalRepo.GetFollow(dstUID, srcUID)
	if err != nil && err != repository.ErrNotFound {
		a.log.Errorf("repo.GetFollow failed %v", err)
		return status.Error(codes.Internal, err.Error())
	}

	a.relationSvc.Unfollow(srcRelation, dstRelation)

	err = a.personalRepo.Save(srcRelation)
	if err != nil {
		a.log.Errorf("personalRepo.Save failed %v", err)
		return status.Error(codes.Internal, err.Error())
	}
	if dstRelation != nil {
		err = a.personalRepo.Save(dstRelation)
		if err != nil {
			a.log.Errorf("personalRepo.Save failed %v", err)
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

func (a *AppService) Block(srcUID, dstUID int64) error {
	var (
		err      error
		relation *entity.PersonalRelation
	)
	relation, err = a.personalRepo.Get(srcUID, dstUID)
	if err != nil && err != repository.ErrNotFound {
		a.log.Errorf("repo.Get failed [%v]", err)
		return status.Error(codes.Internal, err.Error())
	}

	if err != nil && err == repository.ErrNotFound {
		relation = entity.NewPersonalRelation(srcUID, dstUID)
	}

	relation.Block()
	err = a.personalRepo.Save(relation)
	if err != nil {
		a.log.Errorf("repo save failed [%v]", err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (a *AppService) Unblock(srcUID, dstUID int64) error {
	var (
		err      error
		relation *entity.PersonalRelation
	)
	relation, err = a.personalRepo.GetBlock(srcUID, dstUID)
	if err != nil && err != repository.ErrNotFound {
		a.log.Errorf("repo.GetBlock failed [%v]", err)
		return status.Error(codes.Internal, err.Error())
	}
	if err != nil && err == repository.ErrNotFound {
		return status.Error(codes.NotFound, "not found data")
	}

	relation.Unblock()
	err = a.personalRepo.Save(relation)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (a *AppService) JoinGroup(srcUID int64, groupID int64) error {
	var (
		err      error
		relation *entity.GroupRelation
	)
	relation = entity.NewGroupRelation(srcUID, groupID)
	err = a.groupRelationRepo.Save(relation)
	if err != nil {
		a.log.Errorf("repo.save failed")
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (a *AppService) QuitGroup(srcUID int64, groupID int64) error {
	var (
		err      error
		relation *entity.GroupRelation
	)

	relation, err = a.groupRelationRepo.Get(srcUID, groupID)
	if err != nil && err != repository.ErrNotFound {
		a.log.Errorf("repo.Get failed [%v]", err)
		return status.Error(codes.Internal, err.Error())
	}
	relation.Quit()
	err = a.groupRelationRepo.Save(relation)
	if err != nil {
		a.log.Errorf("repo.Save failed [%v]", err)
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (a *AppService) CreateGroup(srcUID int64, name string) (group *entity.Group, err error) {
	var (
		groupRelation *entity.GroupRelation
	)
	group, groupRelation, err = a.relationSvc.CreateGroup(srcUID, name, a.groupRepo.GetSeqName())
	if err != nil {
		a.log.Errorf("create group failed [%v]", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = a.groupRepo.Save(group)
	if err != nil {
		a.log.Errorf("repo.save failed")
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = a.groupRelationRepo.Save(groupRelation)
	if err != nil {
		a.log.Errorf("repo.save failed")
		return nil, status.Error(codes.Internal, err.Error())
	}
	return
}

func (a *AppService) GetGroup(groupID int64) (group *entity.Group, err error) {
	return a.groupRepo.Get(groupID)
}

func (a *AppService) ChangeGroupInfo(uid int64, group *entity.Group) (err error) {

	err = a.relationSvc.ChangeGroup(uid, group, a.groupRelationRepo)
	if err != nil {
		if errors.Cause(err) == repository.ErrNotFound {
			return status.Error(codes.NotFound, "not found group")
		}

		if errors.Cause(err) == service.ErrNoPermission {
			return status.Error(codes.PermissionDenied, "not have permission")
		}

		return status.Error(codes.Internal, err.Error())

	}

	err = a.groupRepo.Save(group)
	if err != nil {
		a.log.Error("repo save failed [%v]", err)
		return status.Error(codes.Internal, err.Error())
	}
	return
}

func (a *AppService) UpgradeGroup(uid, groupID int64, typ int32) error {
	err := a.relationSvc.UpgradeGroup(uid, groupID, typ, a.groupRepo, a.groupRelationRepo)
	switch errors.Cause(err) {
	case service.ErrNoPermission:
		return status.Error(codes.PermissionDenied, "not have permission")
	case service.ErrParameterIncorrect:
		return status.Error(codes.InvalidArgument, "group type not correct")
	default:
		a.log.Errorf("upgrade group failed  [%v]", err)
		return status.Error(codes.Internal, err.Error())
	}

}
