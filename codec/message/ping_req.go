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
func (m *PingReq) Decode(data []byte) error {
	return nil
}

// EncodeTo EncodeTo
func (m *PingReq) EncodeTo(bw *bufio.Writer) error {
	if _, err := bw.Write([]byte{0xc0, 0x0}); err != nil {
		return err
	}
	return nil
}
