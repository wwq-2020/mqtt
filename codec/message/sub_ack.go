package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	Register(controltype.SubAck, NewSubAck)
}

// SubAck SubAck
type SubAck struct {
	flags uint8
}

// NewSubAck NewSubAck
func NewSubAck(flags uint8) Message {
	return &SubAck{
		flags: flags,
	}
}

// Decode Decode
func (m *SubAck) Decode(data []byte) error {
	return nil
}

// EncodeTo EncodeTo
func (m *SubAck) EncodeTo(bw *bufio.Writer) error {
	return nil
}
