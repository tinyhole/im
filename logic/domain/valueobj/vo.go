package valueobj

const (
	DeviceTypeUnknown = 0
	PCLinux           = 1
	PCWindows         = 2
	PhoneAndroid      = 3
	PhoneIOS          = 4
	PadAndroid        = 5
	PadIOS            = 6
)

type SessionInfo struct {
	UID       int64  `json:"uid"`
	DeviceTye int32  `json:"device_type"`
	RemoteIP  string `json:"remote_ip"`
	ApID      int32  `json:"ap_id"`
	ApName    string `json:"ap_name"`
	ApFid     int64  `json:"file_id"`
	UpdateAt  int64  `json:"update_at"`
}
