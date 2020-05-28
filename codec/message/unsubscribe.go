package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	Register(controltype.UnSubscribe, NewUnSubscribe)
}

// UnSubscribe UnSubscribe
type UnSubscribe struct {
	flags uint8
}

// NewUnSubscribe NewUnSubscribe
func NewUnSubscribe(flags uint8) Message {
	return &UnSubscribe{
		flags: flags,
	}
}

// Decode Decode
func (m *UnSubscribe) Decode(data []byte) error {
	return nil
}

// EncodeTo EncodeTo
func (m *UnSubscribe) EncodeTo(bw *bufio.Writer) error {
	return nil
}
