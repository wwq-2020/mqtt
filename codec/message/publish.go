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
	Register(controltype.Publish, NewPublish)
}

// Publish Publish
type Publish struct {
	flags   uint8
	Dup     bool
	Retain  bool
	Qos     qos.Qos
	Topic   string
	Payload []byte
	MsgID   uint16
}

// NewPublish NewPublish
func NewPublish(flags uint8) Message {
	return &Publish{
		flags:  flags,
		Dup:    flags&0x08 > 0,
		Qos:    qos.ParseQos(flags & 0x06 >> 1),
		Retain: flags&0x01 > 0,
	}
}

// Decode Decode
func (m *Publish) Decode(data []byte) error {
	b := bytesutil.Get(data)
	defer bytesutil.Put(b)
	topicLen := b.ReadUint16()
	m.Topic = b.ReadStr(topicLen)
	if m.hasMsgID() {
		m.MsgID = b.ReadUint16()
	}
	m.Payload = b.ReadAll()
	return nil
}

func (m *Publish) hasMsgID() bool {
	return m.Qos == qos.AtLeastOnceDelivery || m.Qos == qos.ExactlyOnceDelivery
}

// EncodeTo EncodeTo
func (m *Publish) EncodeTo(bw *bufio.Writer) error {
	if err := m.writeHeader(bw); err != nil {
		return err
	}
	if err := m.writeBody(bw); err != nil {
		return err
	}
	return nil
}

func (m *Publish) writeHeader(bw *bufio.Writer) error {
	h := header.Get()
	defer header.Put(h)
	header.SetControlType(controltype.Publish)
	m.setFlags()
	header.SetFlags(m.flags)
	if err := header.EncodeTo(bw); err != nil {
		return err
	}
	return nil
}

func (m *Publish) writeBody(bw *bufio.Writer) error {
	dataLen := 2 + uint32(len(m.Topic)) + uint32(len(m.Payload))
	if m.hasMsgID() {
		dataLen += 2
	}
	bytesBuffer := bytespool.Get(int(dataLen))
	defer bytespool.Put(bytesBuffer)
	for dataLen/0x80 > 0 {
		mod := dataLen % 0x80
		if err := bytesBuffer.WriteByte(byte(mod)); err != nil {
			return err
		}
	}
	bytesBuffer.WriteString(m.Topic)
	if m.hasMsgID() {
		if _, err := bytesBuffer.Write(util.Uint16ToBytes(m.MsgID)); err != nil {
			return err
		}
	}
	if _, err := bytesBuffer.Write(m.Payload); err != nil {
		return err
	}
	if _, err := bw.Write(bytesBuffer.Bytes()); err != nil {
		return err
	}
	return nil
}

func (m *Publish) setFlags() {
	m.flags |= util.BoolToUint8(m.Dup) << 3
	m.flags |= util.BoolToUint8(m.Dup)
	m.flags |= m.Qos.ToUint8() << 1
}
