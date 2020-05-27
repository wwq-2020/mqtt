package message

import (
	"bufio"
	"fmt"

	"github.com/wwq1988/mqtt/codec/controltype"
)

var (
	type2Factory = make(map[controltype.Type]Factory)
)

// Message Message
type Message interface {
	Decode(data []byte) error
	Encode(bw *bufio.Writer) error
}

// Factory Factory
type Factory func(flags uint8) Message

// Register Register
func Register(t controltype.Type, factory Factory) {
	factory, exist := type2Factory[t]
	if exist {
		panic(fmt.Sprintf("dup Register for type:%d", t))
	}
	type2Factory[t] = factory
}

// GetFactory GetFactory
func GetFactory(t controltype.Type) Factory {
	return type2Factory[t]
}
