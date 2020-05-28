package util

import "encoding/binary"

// BoolToUint8 BoolToUint8
func BoolToUint8(src bool) uint8 {
	if src {
		return 0x01
	}
	return 0x00
}

// Uint16ToBytes Uint16ToBytes
func Uint16ToBytes(src uint16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, src)
	return buf
}
