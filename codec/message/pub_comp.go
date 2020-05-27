package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	Register(controltype.PubComp, NewPubComp)
}

// PubComp PubComp
type PubComp struct {
	flags uint8
}

// NewPubComp NewPubComp
func NewPubComp(flags uint8) Message {
	return &PubComp{
		flags: flags,
	}
}

// Decode Decode
func (m *PubComp) Decode(data []byte) error {
	return nil
}

// Encode Encode
func (m *PubComp) Encode(bw *bufio.Writer) error {
	return nil
}
