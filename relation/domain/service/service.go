package service

import (
	"github.com/pkg/errors"
	"github.com/tinyhole/im/relation/domain/entity"
	"github.com/tinyhole/im/relation/domain/gateway"
	"github.com/tinyhole/im/relation/domain/repository"
)

var (
	ErrSourceRelationIsNil    = errors.New("source relation is nil")
	ErrGroupRelationIncorrect = errors.New("group relation incorrect")
	ErrNoPermission           = errors.New("no permission")
	ErrParameterIncorrect     = errors.New("parameter incorrect")
)

type RelationService interface {
	Follow(srcRelation, dstRelation *entity.PersonalRelation) error
	Unfollow(srcRelation, dstRelation *entity.PersonalRelation) error
	//Block(srcRelation, dstRelation *entity.PersonalRelation)(error)
	CreateGroup(srcUID int64, groupName, sequenceName string) (*entity.Group, *entity.GroupRelation, error)
	ChangeGroup(srcUID int64, group *entity.Group, repo repository.GroupRelation) error
	UpgradeGroup(uid, groupID int64, typ int32, groupRepo repository.Group, gr repository.GroupRelation) error
}

type relationService struct {
	idCli gateway.SequenceClient
}

func NewRelationService(idCli gateway.SequenceClient) RelationService {
	return &relationService{
		idCli: idCli,
	}
}

//Follow 关注某人
//srcRelation 关注者与被关注者 在关注者方的关系记录
//dstRelation 被关注者和关注者 在关注者方的关系记录
func (r *relationService) Follow(srcRelation, dstRelation *entity.PersonalRelation) error {
	if srcRelation == nil {
		return ErrSourceRelationIsNil
	}

	srcRelation.Follow()
	if dstRelation == nil {
		return nil
	}

	//两者相互关注自动成为好友
	if dstRelation.IsFollow == true {
		srcRelation.Friend()
		dstRelation.Friend()
	}
	return nil
}

func (r *relationService) Unfollow(srcRelation, dstRelation *entity.PersonalRelation) error {
	if srcRelation == nil {
		return ErrSourceRelationIsNil
	}
	srcRelation.Unfollow()
	//单方面取消关注
	if dstRelation != nil && dstRelation.IsFriend == true {
		dstRelation.Unfriend()
		return nil
	}
	return nil
}

func (r *relationService) CreateGroup(srcUID int64, groupName, sequenceName string) (group *entity.Group,
	groupRelation *entity.GroupRelation, err error) {
	var (
		id int64
	)
	group = entity.NewGroup(groupName)
	id, err = r.idCli.GenerateGroupID(sequenceName)
	if err != nil {
		err = errors.Wrap(err, "idCli.GetSeq failed")
		return
	}
	group.ID = id

	groupRelation = entity.NewGroupRelation(srcUID, id)
	groupRelation.Join()
	groupRelation.ChangeRole(entity.GR_Owner)

	return
}

func (r *relationService) ChangeGroup(uid int64, group *entity.Group, repo repository.GroupRelation) error {
	var (
		err           error
		groupRelation *entity.GroupRelation
	)

	groupRelation, err = repo.Get(uid, group.ID)
	if err != nil && err != repository.ErrNotFound {
		return errors.Wrap(err, "repo error")
	}

	if err != nil && err == repository.ErrNotFound {
		return ErrGroupRelationIncorrect
	}

	if groupRelation.Role != entity.GR_Owner ||
		groupRelation.Role != entity.GR_Manager {
		return ErrNoPermission
	}
	return nil
}

func (r *relationService) UpgradeGroup(uid, groupID int64, typ int32,
	groupRepo repository.Group, grRepo repository.GroupRelation) (err error) {
	var (
		group  *entity.Group
		newTyp int32
	)

	_, err = grRepo.Get(uid, groupID)
	if err != nil {
		err = ErrNoPermission
		return
	}
	group, err = groupRepo.Get(groupID)
	if err != nil {
		return err
	}
	switch typ {
	case entity.GT_General:
	case entity.GT_Super:
	case entity.GT_Temporary:
		newTyp = typ
	default:
		return ErrParameterIncorrect
	}

	group.Type = newTyp
	groupRepo.Save(group)

	return nil
}
