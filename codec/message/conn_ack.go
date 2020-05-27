package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	Register(controltype.ConnAck, NewConnAck)
}

// ConnAck ConnAck
type ConnAck struct {
	flags uint8
}

// NewConnAck NewConnAck
func NewConnAck(flags uint8) Message {
	return &ConnAck{
		flags: flags,
	}
}

// Decode Decode
func (m *ConnAck) Decode(br *bufio.Reader) error {
	return nil
}

// Encode Encode
func (m *ConnAck) Encode(br *bufio.Writer) error {
	return nil
}
