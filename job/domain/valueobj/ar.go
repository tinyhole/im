package valueobj

import "container/list"

type MessageNotify struct {
	SessMap map[int32]*list.List
	Notify  *Notify
}
