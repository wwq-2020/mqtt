package bytes

import (
	"encoding/binary"
	"fmt"
)

// ToUInt16 ToUInt16
func ToUInt16(bytes []byte) uint16 {
	if len(bytes) != 2 {
		panic(fmt.Sprintf("unexpected bytes len:%d", len(bytes)))
	}
	return binary.BigEndian.Uint16(bytes)
}

// ToUInt32 ToUInt32
func ToUInt32(bytes []byte) uint32 {
	if len(bytes) != 2 {
		panic(fmt.Sprintf("unexpected bytes len:%d", len(bytes)))
	}
	return binary.BigEndian.Uint32(bytes)
}
