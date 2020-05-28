package message

import (
	"bufio"
	"errors"
)

// UnKnown UnKnown
type UnKnown struct {
	flags uint8
}

// NewUnKnown NewUnKnown
func NewUnKnown(flags uint8) Message {
	return &UnKnown{
		flags: flags,
	}
}

// Decode Decode
func (m *UnKnown) Decode(data []byte) error {
	return errors.New("unknown type")
}

// EncodeTo EncodeTo
func (m *UnKnown) EncodeTo(bw *bufio.Writer) error {
	return errors.New("unknown type")
}
