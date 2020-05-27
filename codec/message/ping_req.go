package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	Register(controltype.PingReq, NewPingReq)
}

type PingReq struct {
	flags uint8
}

// NewPingReq NewPingReq
func NewPingReq(flags uint8) Message {
	return &PingReq{
		flags: flags,
	}
}

// Decode Decode
func (m *PingReq) Decode(br *bufio.Reader) error {
	return nil
}

// Encode Encode
func (m *PingReq) Encode(bw *bufio.Writer) error {
	return nil
}
