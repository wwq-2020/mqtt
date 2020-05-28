package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	Register(controltype.Connnect, NewConnect)
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
func (m *Connect) Decode(data []byte) error {
	return nil
}

// EncodeTo EncodeTo
func (m *Connect) EncodeTo(bw *bufio.Writer) error {
	return nil
}
