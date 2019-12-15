package gateway

type RelationClient interface {
	GetGroupType(groupID int64) (typ int32, err error)
	ListGroupMember(groupID int64, page, pageSize int32) (rets []int64, total int32, err error)
}
