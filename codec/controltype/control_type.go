package controltype

// Type 控制包类型
type Type uint8

// ControlPacketTypes
const (
	Reserved    Type = 0
	Connnect    Type = 1
	ConnAck     Type = 2
	Publish     Type = 3
	PublishAck  Type = 4
	PubRec      Type = 5
	PubRel      Type = 6
	PubComp     Type = 7
	Subscribe   Type = 8
	SubAck      Type = 9
	UnSubscribe Type = 10
	UnSubAck    Type = 11
	PingReq     Type = 12
	PingResp    Type = 13
	Disconnect  Type = 14
	Auth        Type = 15
)
