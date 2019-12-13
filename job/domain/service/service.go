package service

import (
	"github.com/pkg/errors"
	"github.com/tinyhole/im/idl/mua/im"
	"github.com/tinyhole/im/job/domain/entity"
	"github.com/tinyhole/im/job/domain/gateway"
	"github.com/tinyhole/im/job/domain/repository"
	"github.com/tinyhole/im/job/domain/util"
	"github.com/tinyhole/im/job/domain/valueobj"
	"time"
)

type JobService struct {
	notifyClient   gateway.ApClient
	relationClient gateway.RelationClient
	sequenceClient gateway.SequenceClient
	sessionClient  gateway.SessionClient
	msgIDCli       util.MsgIDClient
	repo           repository.Inbox
}

func NewJobService(notifyClient gateway.ApClient,
	relationClient gateway.RelationClient,
	sessionClient gateway.SessionClient,
	sequenceClient gateway.SequenceClient,
	msgIDClient util.MsgIDClient,
	repo repository.Inbox) *JobService {
	return &JobService{
		notifyClient:   notifyClient,
		relationClient: relationClient,
		sequenceClient: sequenceClient,
		sessionClient:  sessionClient,
		msgIDCli:       msgIDClient,
		repo:           repo,
	}
}

func (j *JobService) ProcessMsg(data *entity.Message) error {
	//消息逻辑
	data.Time = time.Now().Unix()
	data.MsgID = j.msgIDCli.GetMsgID()
	if data.MsgType == int32(im.MsgType_MsgTypePrivte) {
		j.PrivateChat(data)
	}

	return nil
}

func (j *JobService) unicastNotify(uid int64, notify *valueobj.MsgNotify) error {
	sessionInfo, err := j.sessionClient.ListSessionInfo(uid)
	if err != nil {
		return errors.WithStack(err)
	}
	if len(sessionInfo) > 0 {
		err = j.notifyClient.PushMsg(sessionInfo[0].ApID, sessionInfo[0].Fid, notify)
	}

	return err
}

func (j *JobService) broadcastNotify(uids []int64, notify *valueobj.MsgNotify) error {
	 j.sessionClient.BatchListSessionInfo(uids)
	return nil
}

func (j *JobService) PrivateChat(data *entity.Message) error {
	seq, err := j.sequenceClient.GetPrivateSeq(data.InboxID)
	if err != nil {
		return errors.WithStack(err)
	}
	data.Seq = seq
	//投递消息



	return nil
}

func (j *JobService) SyncPublicInboxMsg(uid int64) ([]*entity.Message, error) {

	return nil, nil
}
