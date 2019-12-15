package gateway

type ApClient interface {
	Unicast(apID int32, fid int64, data []byte) error
	Broadcast(apID int32, fid []int64, data []byte) error
}
