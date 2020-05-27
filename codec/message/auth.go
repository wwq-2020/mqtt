package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

func init() {
	Register(controltype.Auth, NewAuth)
}

// Auth Auth
type Auth struct {
	flags uint8
}

// NewAuth NewAuth
func NewAuth(flags uint8) Message {
	return &Auth{
		flags: flags,
	}
}

// Decode Decode
func (m *Auth) Decode(br *bufio.Reader) error {
	return nil
}

// Encode Encode
func (m *Auth) Encode(bw *bufio.Writer) error {
	return nil
}
