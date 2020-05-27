package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	register(controltype.Connnect, NewConnect)
}

// Connect Connect
type Connect struct {
	flags uint8
}

// NewConnect NewConnect
func NewConnect(flags uint8) Message {
	return &Connect{
		flags: flags,
	}
}

// Decode Decode
func (m *Connect) Decode(br *bufio.Reader) error {
	return nil
}

// Encode Encode
func (m *Connect) Encode(bw *bufio.Writer) error {
	return nil
}
