package codec

import (
	"bufio"
	"fmt"
	"io"

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
func New(rw io.ReadWriter) Codec {
	return &codec{
		br: bufio.NewReader(rw),
		bw: bufio.NewWriter(rw),
	}
}

func (c *codec) Decode() (message.Message, error) {
	header := header.New()
	if err := header.Decode(c.br); err != nil {
		return nil, err
	}
	messageFactory := message.GetFactory(header.ControlType())
	message := messageFactory(header.Flags())
	buf := make([]byte, header.DataLen())
	if _, err := io.ReadFull(c.br, buf); err != nil {
		return nil, err
	}
	if err := message.Decode(buf); err != nil {
		return nil, fmt.Errorf("got err:%#v for type:%d", err, header.ControlType())
	}
	return message, nil
}

func (c *codec) Encode(message message.Message) error {
	if err := message.Encode(c.bw); err != nil {
		return err
	}
	return nil
}
