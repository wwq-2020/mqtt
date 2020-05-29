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
	Register(controltype.PubRec, NewPubRec)
}

// PubRec PubRec
type PubRec struct {
	flags uint8
	MsgID uint16
}

// NewPubRec NewPubRec
func NewPubRec(flags uint8) Message {
	return &PubRec{
		flags: flags,
	}
}

// Decode Decode
func (m *PubRec) Decode(data []byte) error {
	b := bytesutil.Get(data)
	defer bytesutil.Put(b)
	m.MsgID = b.ReadUint16()
	return nil
}

// EncodeTo EncodeTo
func (m *PubRec) EncodeTo(bw *bufio.Writer) error {
	if err := m.writeHeader(bw); err != nil {
		return err
	}
	if err := m.writeBody(bw); err != nil {
		return err
	}
	return nil
}

func (m *PubRec) writeHeader(bw *bufio.Writer) error {
	h := header.Get()
	defer header.Put(h)
	h.SetControlType(controltype.PubRel)
	m.setFlags()
	h.SetFlags(m.flags)
	if err := h.EncodeTo(bw); err != nil {
		return err
	}
	return nil
}

func (m *PubRec) writeBody(bw *bufio.Writer) error {
	dataLen := 2

	bytesBuffer := bytespool.Get(int(dataLen))
	defer bytespool.Put(bytesBuffer)

	bytesBuffer.Write(util.Uint16ToBytes(m.MsgID))

	if _, err := bw.Write(bytesBuffer.Bytes()); err != nil {
		return err
	}
	return nil
}

func (m *PubRec) setFlags() {

}
