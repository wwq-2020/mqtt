package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	Register(controltype.Publish, NewPublish)
}

// Publish Publish
type Publish struct {
	flags uint8
}

// NewPublish NewPublish
func NewPublish(flags uint8) Message {
	return &Publish{
		flags: flags,
	}
}

// Decode Decode
func (m *Publish) Decode(data []byte) error {
	return nil
}

// EncodeTo EncodeTo
func (m *Publish) EncodeTo(bw *bufio.Writer) error {
	return nil
}
