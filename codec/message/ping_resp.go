package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	Register(controltype.PingResp, NewPingResp)
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
func (m *PingResp) Decode(data []byte) error {
	return nil
}

// EncodeTo EncodeTo
func (m *PingResp) EncodeTo(bw *bufio.Writer) error {
	if _, err := bw.Write([]byte{0xd0, 0x0}); err != nil {
		return err
	}
	return nil
}
