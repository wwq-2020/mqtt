package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	Register(controltype.Disconnect, NewDisconnect)
}

// DisConnnect DisConnnect
type DisConnnect struct {
	flags uint8
}

// NewDisconnect NewDisconnect
func NewDisconnect(flags uint8) Message {
	return &Connect{
		flags: flags,
	}
}

// Decode Decode
func (m *DisConnnect) Decode(data []byte) error {
	return nil
}

// EncodeTo EncodeTo
func (m *DisConnnect) EncodeTo(bw *bufio.Writer) error {
	if _, err := bw.Write([]byte{0xe0, 0x0}); err != nil {
		return err
	}
	return nil
}
