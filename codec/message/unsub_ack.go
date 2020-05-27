package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	register(controltype.UnSubAck, NewUnSubAck)
}

// UnSubAck UnSubAck
type UnSubAck struct {
	flags uint8
}

// NewUnSubAck NewUnSubAck
func NewUnSubAck(flags uint8) Message {
	return &UnSubAck{
		flags: flags,
	}
}

// Decode Decode
func (m *UnSubAck) Decode(br *bufio.Reader) error {
	return nil
}

// Encode Encode
func (m *UnSubAck) Encode(bw *bufio.Writer) error {
	return nil
}
