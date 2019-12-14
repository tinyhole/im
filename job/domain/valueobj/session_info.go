package valueobj

/*
   DeviceTypeUnknown = 0;
   PCLinux = 1;
   PCWindows = 2;
   PhoneAndroid = 3;
   PhoneIOS = 4;
   PadAndroid = 5;
   PadIOS = 6;
*/

type SessionInfo struct {
	Uid        int64
	Fid        int64
	ApID       int32
}
