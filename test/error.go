package main

type RspError struct {
	ErrCode int32
	ErrText string
}

func (r *RspError) Error() string {
	return r.ErrText
}
