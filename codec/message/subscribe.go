package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	register(controltype.Subscribe, NewSubscribe)
}

// Subscribe Subscribe
type Subscribe struct {
	flags uint8
}

// NewSubscribe NewSubscribe
func NewSubscribe(flags uint8) Message {
	return &Subscribe{
		flags: flags,
	}
}

// Decode Decode
func (m *Subscribe) Decode(br *bufio.Reader) error {
	return nil
}

// Encode Encode
func (m *Subscribe) Encode(bw *bufio.Writer) error {
	return nil
}
