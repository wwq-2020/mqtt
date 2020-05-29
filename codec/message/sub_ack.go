package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/bytespool"
	"github.com/wwq1988/mqtt/bytesutil"
	"github.com/wwq1988/mqtt/codec/controltype"
	"github.com/wwq1988/mqtt/codec/header"
	"github.com/wwq1988/mqtt/codec/qos"
	"github.com/wwq1988/mqtt/util"
)

func init() {
	Register(controltype.SubAck, NewSubAck)
}

// SubAck SubAck
type SubAck struct {
	flags uint8
	MsgID uint16
	Qoses []qos.Qos
}

// NewSubAck NewSubAck
func NewSubAck(flags uint8) Message {
	return &SubAck{
		flags: flags,
	}
}

// Decode Decode
func (m *SubAck) Decode(data []byte) error {
	b := bytesutil.Get(data)
	defer bytesutil.Put(b)
	m.MsgID = b.ReadUint16()
	for !b.EOF() {
		m.Qoses = append(m.Qoses, qos.Parse(b.ReadByte()))
	}
	return nil
}

// EncodeTo EncodeTo
func (m *SubAck) EncodeTo(bw *bufio.Writer) error {
	if err := m.writeHeader(bw); err != nil {
		return err
	}
	if err := m.writeBody(bw); err != nil {
		return err
	}
	return nil
}

func (m *SubAck) writeHeader(bw *bufio.Writer) error {
	h := header.Get()
	defer header.Put(h)
	h.SetControlType(controltype.SubAck)
	m.setFlags()
	h.SetFlags(m.flags)
	if err := h.EncodeTo(bw); err != nil {
		return err
	}
	return nil
}

func (m *SubAck) writeBody(bw *bufio.Writer) error {
	dataLen := 2 + uint32(len(m.Qoses))

	bytesBuffer := bytespool.Get(int(dataLen))
	defer bytespool.Put(bytesBuffer)

	bytesBuffer.Write(util.Uint16ToBytes(m.MsgID))
	for _, qos := range m.Qoses {
		bytesBuffer.WriteByte(byte(qos.ToUint8()))
	}

	if _, err := bw.Write(bytesBuffer.Bytes()); err != nil {
		return err
	}
	return nil
}

func (m *SubAck) setFlags() {

}
