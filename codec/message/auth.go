package message

import (
	"bufio"
)

func init() {
	// Register(controltype.Auth, NewAuth)
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
func (m *Auth) Decode(data []byte) error {
	return nil
}

// EncodeTo EncodeTo
func (m *Auth) EncodeTo(bw *bufio.Writer) error {
	return nil
}
