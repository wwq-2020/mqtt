package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	Register(controltype.PubRec, NewPubRec)
}

// PubRec PubRec
type PubRec struct {
	flags uint8
}

// NewPubRec NewPubRec
func NewPubRec(flags uint8) Message {
	return &PubRec{
		flags: flags,
	}
}

// Decode Decode
func (m *PubRec) Decode(br *bufio.Reader) error {
	return nil
}

// Encode Encode
func (m *PubRec) Encode(bw *bufio.Writer) error {
	return nil
}
