package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	Register(controltype.PubRel, NewPubRel)
}

// PubRel PubRel
type PubRel struct {
	flags uint8
}

// NewPubRel NewPubRel
func NewPubRel(flags uint8) Message {
	return &PubRel{
		flags: flags,
	}
}

// Decode Decode
func (m *PubRel) Decode(data []byte) error {
	return nil
}

// EncodeTo EncodeTo
func (m *PubRel) EncodeTo(bw *bufio.Writer) error {
	return nil
}
