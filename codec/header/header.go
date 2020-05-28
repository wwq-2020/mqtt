package header

import (
	"bufio"

	"github.com/wwq1988/mqtt/codec/controltype"
)

// Header Header
type Header interface {
	Flags() uint8
	SetFlags(uint8)
	DataLen() uint32
	ControlType() controltype.Type
	SetControlType(controltype.Type)
	EncodeTo(*bufio.Writer) error
	Decode(*bufio.Reader) error
}

type header struct {
	flags       uint8
	dataLen     uint32
	controlType controltype.Type
}

// New New
func New() Header {
	return &header{}
}

func (h *header) Flags() uint8 {
	return h.flags
}

func (h *header) SetFlags(flags uint8) {
	h.flags = flags
}

func (h *header) DataLen() uint32 {
	return h.dataLen
}

func (h *header) ControlType() controltype.Type {
	return h.controlType
}

func (h *header) SetControlType(controlType controltype.Type) {
	h.controlType = controlType
}

func (h *header) Decode(br *bufio.Reader) error {
	b, err := br.ReadByte()
	if err != nil {
		return err
	}
	h.flags = b & 0x0f

	multiplier := uint32(1)
	digit := byte(0x80)
	for (digit & 0x80) != 0 {
		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		digit = b
		h.dataLen += uint32(digit&0x7f) * multiplier
		multiplier *= 128
	}
	return nil
}

func (h *header) EncodeTo(bw *bufio.Writer) error {
	b := (0xf0 & uint8(h.controlType) << 4) | (0x0f & h.flags)
	if err := bw.WriteByte(byte(b)); err != nil {
		return err
	}

	return nil
}
