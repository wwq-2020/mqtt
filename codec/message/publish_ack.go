package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	Register(controltype.PublishAck, NewPublishAck)
}

// PublishAck PublishAck
type PublishAck struct {
	flags uint8
}

// NewPublishAck NewPublishAck
func NewPublishAck(flags uint8) Message {
	return &PublishAck{
		flags: flags,
	}
}

// Decode Decode
func (m *PublishAck) Decode(data []byte) error {
	return nil
}

// EncodeTo EncodeTo
func (m *PublishAck) EncodeTo(bw *bufio.Writer) error {
	return nil
}
