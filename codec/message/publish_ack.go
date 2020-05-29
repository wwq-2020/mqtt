package message

import (
	"bufio"

	"github.com/wwq1988/mqtt/bytespool"
	"github.com/wwq1988/mqtt/bytesutil"
	"github.com/wwq1988/mqtt/codec/controltype"
	"github.com/wwq1988/mqtt/codec/header"
	"github.com/wwq1988/mqtt/util"
)

func init() {
	Register(controltype.PublishAck, NewPublishAck)
}

// PublishAck PublishAck
type PublishAck struct {
	flags uint8
	MsgID uint16
}

// NewPublishAck NewPublishAck
func NewPublishAck(flags uint8) Message {
	return &PublishAck{
		flags: flags,
	}
}

// Decode Decode
func (m *PublishAck) Decode(data []byte) error {
	b := bytesutil.Get(data)
	defer bytesutil.Put(b)
	m.MsgID = b.ReadUint16()
	return nil
}

// EncodeTo EncodeTo
func (m *PublishAck) EncodeTo(bw *bufio.Writer) error {
	h := header.Get()
	defer header.Put(h)
	h.SetControlType(controltype.Publish)
	m.setFlags()
	h.SetFlags(m.flags)
	if err := h.EncodeTo(bw); err != nil {
		return err
	}
	return nil
}

func (m *PublishAck) writeHeader(bw *bufio.Writer) error {
	h := header.Get()
	defer header.Put(h)
	h.SetControlType(controltype.Publish)
	m.setFlags()
	h.SetFlags(m.flags)
	if err := h.EncodeTo(bw); err != nil {
		return err
	}
	return nil
}

func (m *PublishAck) writeBody(bw *bufio.Writer) error {
	dataLen := uint16(2)
	bytesBuffer := bytespool.Get(int(dataLen))
	defer bytespool.Put(bytesBuffer)
	bytesBuffer.WriteByte(byte(dataLen))
	bytesBuffer.Write(util.Uint16ToBytes(m.MsgID))
	return nil
}

func (m *PublishAck) setFlags() {

}
