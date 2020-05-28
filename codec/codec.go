package codec

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wwq1988/mqtt/bufiopool"
	"github.com/wwq1988/mqtt/codec/header"
	"github.com/wwq1988/mqtt/codec/message"
)

// Codec 编解码器
type Codec interface {
	Encode(message message.Message) error
	Decode() (message.Message, error)
}

type codec struct {
	br *bufio.Reader
	bw *bufio.Writer
}

// New New
func New() Codec {
	return &codec{}
}

func (c *codec) Decode() (message.Message, error) {
	header := header.Get()
	if err := header.Decode(c.br); err != nil {
		return nil, err
	}
	messageFactory := message.GetFactory(header.ControlType())
	message := messageFactory(header.Flags())
	buf, err := c.readBody(header.DataLen())
	if err != nil {
		return nil, err
	}
	if err := message.Decode(buf); err != nil {
		return nil, fmt.Errorf("got err:%#v for type:%d", err, header.ControlType())
	}
	return message, nil
}

func (c *codec) readBody(dataLen uint32) ([]byte, error) {
	if dataLen == 0 {
		return nil, nil
	}
	buf := make([]byte, dataLen)
	if _, err := io.ReadFull(c.br, buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func (c *codec) Encode(message message.Message) error {
	if err := message.EncodeTo(c.bw); err != nil {
		return err
	}
	return nil
}

func (c *codec) Reset(rw io.ReadWriter) {
	c.resetBr(rw)
	c.resetBw(rw)
}

func (c *codec) resetBr(r io.Reader) {
	if c.br != nil {
		c.br.Reset(r)
		return
	}
	c.br = bufiopool.GetReader(r)
}

func (c *codec) resetBw(w io.Writer) {
	if c.bw != nil {
		c.bw.Reset(w)
		return
	}
	c.bw = bufiopool.GetWriter(w)
}
