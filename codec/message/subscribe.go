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
	Register(controltype.Subscribe, NewSubscribe)
}

// Topic Topic
type Topic struct {
	Name string
	Qos  uint8
}

// Subscribe Subscribe
type Subscribe struct {
	flags  uint8
	MsgID  uint16
	Topics []*Topic
}

// NewSubscribe NewSubscribe
func NewSubscribe(flags uint8) Message {
	return &Subscribe{
		flags: flags,
	}
}

// Decode Decode
func (m *Subscribe) Decode(data []byte) error {
	b := bytesutil.Get(data)
	defer bytesutil.Put(b)
	m.MsgID = b.ReadUint16()
	for !b.EOF() {
		dataLen := b.ReadUint16()
		name := b.ReadStr(dataLen)
		qos := uint8(b.ReadByte())
		m.Topics = append(m.Topics, &Topic{
			Name: name,
			Qos:  qos,
		})
	}
	return nil
}

// EncodeTo EncodeTo
func (m *Subscribe) EncodeTo(bw *bufio.Writer) error {
	if err := m.writeHeader(bw); err != nil {
		return err
	}
	if err := m.writeBody(bw); err != nil {
		return err
	}
	return nil
}

func (m *Subscribe) writeHeader(bw *bufio.Writer) error {
	h := header.Get()
	defer header.Put(h)
	h.SetControlType(controltype.Subscribe)
	m.setFlags()
	h.SetFlags(m.flags)
	if err := h.EncodeTo(bw); err != nil {
		return err
	}
	return nil
}

func (m *Subscribe) writeBody(bw *bufio.Writer) error {
	dataLen := uint32(2)
	for _, topic := range m.Topics {
		dataLen += 2 + uint32(len(topic.Name)) + 1
	}
	bytesBuffer := bytespool.Get(int(dataLen))
	defer bytespool.Put(bytesBuffer)

	bytesBuffer.Write(util.Uint16ToBytes(m.MsgID))

	for _, topic := range m.Topics {
		dataLen += 2 + uint32(len(topic.Name)) + 1
		bytesBuffer.Write(util.Uint16ToBytes(uint16(len(topic.Name))))
		bytesBuffer.WriteString(topic.Name)
		bytesBuffer.WriteByte(byte(topic.Qos))
	}

	if _, err := bw.Write(bytesBuffer.Bytes()); err != nil {
		return err
	}
	return nil
}

func (m *Subscribe) setFlags() {

}
