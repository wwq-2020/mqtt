package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	Register(controltype.Subscribe, NewSubscribe)
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
func (m *Subscribe) Decode(data []byte) error {

	return nil
}

// EncodeTo EncodeTo
func (m *Subscribe) EncodeTo(bw *bufio.Writer) error {
	return nil
}
