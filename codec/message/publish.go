package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	register(controltype.Publish, NewPublish)
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
func (m *Publish) Decode(br *bufio.Reader) error {
	return nil
}

// Encode Encode
func (m *Publish) Encode(bw *bufio.Writer) error {
	return nil
}
