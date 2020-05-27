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
func (m *SubAck) Decode(br *bufio.Reader) error {
	return nil
}

// Encode Encode
func (m *SubAck) Encode(bw *bufio.Writer) error {
	return nil
}
