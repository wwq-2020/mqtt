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
	Register(controltype.Connnect, NewConnect)
}

// Connect Connect
type Connect struct {
	flags            uint8
	ProtocolName     string
	Version          uint8
	Flags            uint8 // connect flags
	Keepalive        uint16
	ClientID         string
	WillTopic        string
	WillMessage      []byte
	Username         string
	Password         string
	UsernameFlag     bool
	PasswordFlag     bool
	WillRetainFlag   bool
	WillQos          qos.Qos
	WillFlag         bool
	CleanSessionFlag bool
}

// NewConnect NewConnect
func NewConnect(flags uint8) Message {
	return &Connect{
		flags: flags,
	}
}

// Decode Decode
func (m *Connect) Decode(data []byte) error {
	b := bytesutil.Get(data)
	defer bytesutil.Put(b)
	m.ProtocolName = b.ReadStr(2)
	m.Version = uint8(b.ReadByte())
	flags := uint8(b.ReadByte())
	m.Keepalive = b.ReadUint16()
	m.ClientID = b.ReadStr(2)

	m.UsernameFlag = flags&(1<<7) > 0
	m.PasswordFlag = flags&(1<<6) > 0
	m.WillRetainFlag = flags&(1<<5) > 0
	m.WillQos = qos.Parse(flags & 0x8f)
	m.WillFlag = flags&(1<<2) > 0
	m.CleanSessionFlag = flags&(1<<1) > 0
	if m.WillFlag {
		m.WillTopic = b.ReadVariableStr()
		m.WillMessage = b.ReadVariable()
	}
	if m.UsernameFlag {
		m.Username = b.ReadVariableStr()
	}
	if m.PasswordFlag {
		m.Password = b.ReadVariableStr()
	}
	return nil
}

// EncodeTo EncodeTo
func (m *Connect) EncodeTo(bw *bufio.Writer) error {
	if err := m.writeHeader(bw); err != nil {
		return err
	}
	if err := m.writeBody(bw); err != nil {
		return err
	}
	return nil
}

func (m *Connect) writeHeader(bw *bufio.Writer) error {
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

func (m *Connect) writeBody(bw *bufio.Writer) error {
	dataLen := 2 + uint32(len(m.ProtocolName)) + 1 + 1 + 2 + uint32(len(m.ClientID)) + 2
	if m.WillFlag {
		dataLen += 2 + uint32(len(m.WillTopic))
		dataLen += 2 + uint32(len(m.WillMessage))
	}
	if m.UsernameFlag {
		dataLen += 2 + uint32(len(m.Username))
	}
	if m.PasswordFlag {
		dataLen += 2 + uint32(len(m.Password))
	}

	bytesBuffer := bytespool.Get(int(dataLen))
	defer bytespool.Put(bytesBuffer)
	bytesBuffer.Write(util.Uint16ToBytes(uint16(len(m.ProtocolName))))
	bytesBuffer.WriteString(m.ProtocolName)

	bytesBuffer.WriteByte(m.Flags)
	bytesBuffer.Write(util.Uint16ToBytes(m.Keepalive))

	bytesBuffer.Write(util.Uint16ToBytes(uint16(len(m.ClientID))))
	bytesBuffer.WriteString(m.ClientID)

	if m.WillFlag {
		bytesBuffer.Write(util.Uint16ToBytes(uint16(len(m.WillTopic))))
		bytesBuffer.WriteString(m.WillTopic)
		bytesBuffer.Write(util.Uint16ToBytes(uint16(len(m.WillMessage))))
		bytesBuffer.Write(m.WillMessage)
	}

	if m.UsernameFlag {
		bytesBuffer.Write(util.Uint16ToBytes(uint16(len(m.Username))))
		bytesBuffer.WriteString(m.Username)
	}

	if m.PasswordFlag {
		bytesBuffer.Write(util.Uint16ToBytes(uint16(len(m.Password))))
		bytesBuffer.WriteString(m.Password)
	}
	if _, err := bw.Write(bytesBuffer.Bytes()); err != nil {
		return err
	}
	return nil
}

func (m *Connect) setFlags() {
	m.Flags |= util.BoolToUint8(m.UsernameFlag) << 7
	m.Flags |= util.BoolToUint8(m.PasswordFlag) << 6
	m.Flags |= util.BoolToUint8(m.WillRetainFlag) << 5
	m.Flags |= m.WillQos.ToUint8() << 3
	m.Flags |= util.BoolToUint8(m.WillFlag) << 2
	m.Flags |= util.BoolToUint8(m.CleanSessionFlag) << 1
}
