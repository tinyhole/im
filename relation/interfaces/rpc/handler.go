package rpc

import (
	"context"
	"github.com/tinyhole/im/idl/mua/im/relation"
	"github.com/tinyhole/im/relation/application"
	"github.com/tinyhole/im/relation/domain/entity"
	"github.com/tinyhole/im/relation/interfaces/objconv"
)

type Handler struct {
	appSvc               *application.AppService
	groupConv            *objconv.GroupConv
	groupRelationConv    *objconv.GroupRelationConv
	personalRelationConv *objconv.PersonalRelationConv
}

func NewHandler(appSrv *application.AppService,
	groupConv *objconv.GroupConv,
	groupRelationConv *objconv.GroupRelationConv,
	personalRelationConv *objconv.PersonalRelationConv) *Handler {
	return &Handler{
		appSvc:               appSrv,
		groupConv:            groupConv,
		groupRelationConv:    groupRelationConv,
		personalRelationConv: personalRelationConv,
	}
}

func (h *Handler) Follow(ctx context.Context, req *relation.FollowReq, rsp *relation.FollowRsp) (err error) {
	err = h.appSvc.Follow(req.SrcUID, req.DstUID)
	return
}

func (h *Handler) Unfollow(ctx context.Context, req *relation.UnfollowReq, rsp *relation.UnfollowRsp) (err error) {
	err = h.appSvc.Unfollow(req.SrcUID, req.DstUID)
	return
}

func (h *Handler) Block(ctx context.Context, req *relation.BlockReq, rsp *relation.BlockRsp) (err error) {
	err = h.appSvc.Block(req.SrcUID, req.DstUID)
	return
}

func (h *Handler) Unblock(ctx context.Context, req *relation.UnblockReq, rsp *relation.UnblockRsp) (err error) {
	err = h.appSvc.Unblock(req.SrcUID, req.DstUID)
	return
}

func (h *Handler) GetPersonalRelation(ctx context.Context, req *relation.GetPersonalRelationReq,
	rsp *relation.GetPersonalRelationRsp) (err error) {
	var (
		data *entity.PersonalRelation
	)
	data, err = h.appSvc.GetPersonalRelation(req.SrcUID, req.DstUID)
	if err != nil {
		return
	}
	rsp.Info = h.personalRelationConv.DO2DTO(data)

	return
}

func (h *Handler) ListPersonalRelation(ctx context.Context, req *relation.ListPersonalRelationReq,
	rsp *relation.ListPersonalRelationRsp) (err error) {
	var (
		data  []*entity.PersonalRelation
		total int
	)
	data, total, err = h.appSvc.ListPersonalRelation(req.UID, int32(req.Type), req.Page, req.PageSize)
	if err != nil {
		return
	}
	rsp.Total = int32(total)
	rsp.Infos = h.personalRelationConv.SliceDO2DTO(data)
	return
}

func (h *Handler) CreateGroup(ctx context.Context, req *relation.CreateGroupReq, rsp *relation.CreateGroupRsp) (err error) {
	var (
		data *entity.Group
	)
	data, err = h.appSvc.CreateGroup(req.SrcUID, req.GroupName)
	if err != nil {
		return err
	}
	rsp.Info = h.groupConv.DO2DTO(data)
	return
}

func (h *Handler) GetGroup(ctx context.Context, req *relation.GetGroupReq, rsp *relation.GetGroupRsp) (err error) {
	var (
		data *entity.Group
	)

	data, err = h.appSvc.GetGroup(req.GroupID)
	if err != nil {
		return err
	}
	rsp.Info = h.groupConv.DO2DTO(data)
	return
}

func (h *Handler) ChangeGroupInfo(ctx context.Context, req *relation.ChangeGroupInfoReq, rsp *relation.ChangeGroupInfoRsp) (err error) {
	var (
		data *entity.Group
	)
	data = h.groupConv.DTO2DO(req.Info)
	err = h.appSvc.ChangeGroupInfo(req.Uid, data)
	return
}

func (h *Handler) UpgradeGroup(ctx context.Context, req *relation.UpgradeGroupReq, rsp *relation.UpgradeGroupRsp) (err error) {
	err = h.appSvc.UpgradeGroup(req.Uid, req.GroupID, int32(req.Type))
	return
}

func (h *Handler) JoinGroup(ctx context.Context, req *relation.JoinGroupReq, rsp *relation.JoinGroupRsp) (err error) {
	err = h.appSvc.JoinGroup(req.UID, req.GroupID)
	return err
}

func (h *Handler) QuitGroup(ctx context.Context, req *relation.QuitGroupReq, rsp *relation.QuitGroupRsp) (err error) {
	err = h.appSvc.QuitGroup(req.UID, req.GroupID)
	return
}

func (h *Handler)ListGroupMemberID(ctx context.Context, req *relation.ListGroupMemberIDReq, rsp *relation.ListGroupMemberIDRsp)(err error){
	rets, total, err :=  h.appSvc.ListGroupMember(req.GroupID, req.Page,req.PageSize)
	if err != nil{
		return err
	}

	rsp.Total = int32(total)
	for _, itr := range rets {

		rsp.Uid = append(rsp.Uid,itr.SrcUID)
	}

	return
}