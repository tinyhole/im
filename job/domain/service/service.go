package service

import (
	"container/list"
	"github.com/pkg/errors"
	"github.com/tinyhole/im/idl/mua/im"
	"github.com/tinyhole/im/idl/mua/im/relation"
	"github.com/tinyhole/im/job/domain/entity"
	"github.com/tinyhole/im/job/domain/gateway"
	"github.com/tinyhole/im/job/domain/repository"
	"github.com/tinyhole/im/job/domain/util"
	"github.com/tinyhole/im/job/domain/valueobj"
	"time"
)


var (
	pageSize = int32(500)
)
type JobService struct {
	relationClient gateway.RelationClient
	sequenceClient gateway.SequenceClient
	sessionClient  gateway.SessionClient
	msgIDCli       util.MsgIDClient
	inboxIDClient  util.InboxIDClient
	repo           repository.Inbox
}

func NewJobService(
	relationClient gateway.RelationClient,
	sessionClient gateway.SessionClient,
	sequenceClient gateway.SequenceClient,
	inboxIDClient util.InboxIDClient) *JobService {
	return &JobService{
		relationClient: relationClient,
		sequenceClient: sequenceClient,
		sessionClient:  sessionClient,
		inboxIDClient:  inboxIDClient,
	}
}

func (j *JobService) ProcessMsg(msg *entity.Message) (rets []*entity.Message, notifies []*valueobj.MessageNotify,err error) {
	//消息逻辑
	msg.Time = time.Now().Unix()
	switch msg.MsgType{
	case int32(im.MsgType_MsgTypePrivte):
		rets, notifies, err = j.privateChat(msg)
	case int32(im.MsgType_MsgTypeGroup):
		rets, notifies, err = j.groupChat(msg)
	default:
		return nil,nil, ErrUnknownMsgType
	}

	return

}

func (j *JobService) groupChat(msg *entity.Message)(rets []*entity.Message, notifies []*valueobj.MessageNotify, err error){
	//拿群类型，决定是写扩散还是读扩散
	groupType, err := j.relationClient.GetGroupType(msg.DstID)
	//写扩散
	if groupType != int32(relation.GroupType_Super){
		rets, notifies, err = j.diffuseWrite(msg)
	}else{ //读扩散
		rets, notifies,err = j.diffuseRead(msg)
	}
	return
}

func (j *JobService)diffuseWrite(msg *entity.Message)(rets []*entity.Message,
	notifies []*valueobj.MessageNotify, err error){
	var (
		page int32
		members []int64
		total int32
	)
	page = 1
	for {
		members, total, err = j.relationClient.ListGroupMember(msg.DstID,page, pageSize)
		if err != nil{
			err =errors.Wrapf(ErrProcessMsgFailed,"list group member failed [%v]",err.Error())
			return
		}
		for _, itr := range members {
			inboxID := j.inboxIDClient.GetPersonalInboxID(itr)
			dstMsg := msg.CopyToInbox(inboxID)
			j.fillMessage(dstMsg)
			rets = append(rets, dstMsg)
			notify := j.generatePersonalNotify(itr, dstMsg)
			if notify != nil{
				notifies = append(notifies, notify)
			}
		}

		if page * pageSize >= total{
			break
		}
	}

	return
}

//diffuseRead read 扩散
func (j *JobService)diffuseRead(msg *entity.Message)(rets []*entity.Message,
	notifies []*valueobj.MessageNotify, err error){
	var (
		members []int64
		total int32
	)
	//填充消息
	inboxID := j.inboxIDClient.GetGroupInboxID(msg.DstID)
	msg.InboxID = inboxID
	err = j.fillMessage(msg)
	rets = append(rets, msg)
	if err != nil{
		err = ErrProcessMsgFailed
		return
	}
	notify := &valueobj.Notify{
		InboxID: inboxID,
		Seq:     msg.Seq,
	}
	page := int32(1)
	for {
		//生成消息通知
		members, total, err = j.relationClient.ListGroupMember(msg.DstID, page, pageSize)
		if err != nil{
			err = ErrGenerateNotifyFailed
			return
		}
		notifies = append(notifies,j.generateGroupNotify(members,notify))
		//所有通知都生成完了
		if pageSize * page >= total {
			break
		}
	}
	return
}

func (j *JobService) privateChat(msg *entity.Message)(rets []*entity.Message,notifies []*valueobj.MessageNotify,err error)  {
	//私聊写扩散
	senderInboxID := j.inboxIDClient.GetPersonalInboxID(msg.SrcID)
	senderMsg := msg
	receiverInboxID := j.inboxIDClient.GetPersonalInboxID(msg.DstID)
	msg.InboxID = senderInboxID
	receiverMsg := msg.CopyToInbox(receiverInboxID)

	err = j.fillMessage(senderMsg)
	if err != nil{
		return
	}
	err = j.fillMessage(receiverMsg)
	if err != nil{
		return
	}
	rets = append(rets,senderMsg,receiverMsg)
	senderNotify := j.generatePersonalNotify(msg.SrcID, senderMsg)
	receiverNotify := j.generatePersonalNotify(msg.DstID, receiverMsg)
	if senderNotify != nil{
		notifies = append(notifies, senderNotify)
	}

	if receiverNotify != nil{

		notifies = append(notifies,receiverNotify)
	}
	return
}

func(j *JobService)fillMessage(msg *entity.Message)(error){
	seq, err := j.sequenceClient.GetPrivateSeq(msg.InboxID)
	if err != nil {
		return errors.Wrap(ErrProcessMsgFailed, "GetPrivateSeq failed")
	}
	msg.Seq = seq
	return nil
}


//generatePersonalNotify 生成个人通知
func(j *JobService) generatePersonalNotify(uid int64,msg *entity.Message)*valueobj.MessageNotify {
	var (
		sesMap map[int32]*list.List
	)
	//对会话按照APID分组
	sesMap = make(map[int32]*list.List)
	sess , err := j.sessionClient.ListSessionInfo(uid)
	for _, itr := range sess{
		if v, ok :=  sesMap[itr.ApID];ok{
			v.PushBack(itr)
		}else{
			sesMap[itr.ApID] = list.New()
			sesMap[itr.ApID].PushBack(itr)
		}
	}
	if err != nil{
		return nil
	}
	return &valueobj.MessageNotify{
		SessMap: sesMap,
		Notify: &valueobj.Notify{
			InboxID: msg.InboxID,
			Seq:     msg.Seq,
		},
	}
}


//generateGroupNotify 生成群通知
func (j *JobService)generateGroupNotify(uids []int64, notify *valueobj.Notify)*valueobj.MessageNotify{
	var (
		sessMap map[int32]*list.List
	)
	//对会话按照APID分组
	sessMap = make(map[int32]*list.List)
	sess,err := j.sessionClient.BatchListSessionInfo(uids)
	for _, itr := range sess{
		if v, ok :=  sessMap[itr.ApID];ok{
			v.PushBack(itr)
		}else{
			sessMap[itr.ApID] = list.New()
			sessMap[itr.ApID].PushBack(itr)
		}
	}
	if err != nil{
		return nil
	}
	return &valueobj.MessageNotify{
		SessMap: sessMap,
		Notify: notify,
	}
}

func (j *JobService)SyncPrivateInboxMsg(uid int64, seq int64, page ,pageSize int32)([]*entity.Message,int,error){
	inboxID := j.inboxIDClient.GetPersonalInboxID(uid)
	end := time.Now()
	start := time.Now().Add(-7 * time.Hour * 24)
	rets, total, err :=j.repo.List(inboxID,seq, start.Unix(),end.Unix(),page,pageSize)

	return rets,total,err
}

func (j *JobService) SyncPublicInboxMsg(uid int64) ([]*entity.Message, error) {

	return nil, nil
}
