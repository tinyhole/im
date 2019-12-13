package gateway

type RelationClient interface {
	GetGroupType(groupID int64) (typ int32, err error)
}
