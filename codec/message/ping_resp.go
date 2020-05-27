package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	register(controltype.PingResp, NewPingResp)
}

// PingResp PingResp
type PingResp struct {
	flags uint8
}

// NewPingResp NewPingResp
func NewPingResp(flags uint8) Message {
	return &PingResp{
		flags: flags,
	}
}

// Decode Decode
func (m *PingResp) Decode(br *bufio.Reader) error {
	return nil
}

// Encode Encode
func (m *PingResp) Encode(bw *bufio.Writer) error {
	return nil
}
