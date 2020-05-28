package bufiopool

import (
	"bufio"
	"io"
	"sync"
)

const (
	// ReaderBufferSize ReaderBufferSize
	ReaderBufferSize = 4 * 1024
	// WriterBufferSize WriterBufferSize
	WriterBufferSize = 4 * 1024
)

var (
	readerPool = &sync.Pool{
		New: func() interface{} {
			return bufio.NewReaderSize(nil, ReaderBufferSize)
		},
	}
	writerPool = &sync.Pool{
		New: func() interface{} {
			return bufio.NewWriterSize(nil, WriterBufferSize)
		},
	}
)

// GetReader GetReader
func GetReader(r io.Reader) *bufio.Reader {
	br := readerPool.Get().(*bufio.Reader)
	br.Reset(r)
	return br
}

// PutReader PutReader
func PutReader(br *bufio.Reader) {
	readerPool.Put(br)
}

// GetWriter GetWriter
func GetWriter(w io.Writer) *bufio.Writer {
	bw := writerPool.Get().(*bufio.Writer)
	bw.Reset(w)
	return bw
}

// PutWriter PutWriter
func PutWriter(bw *bufio.Writer) {
	writerPool.Put(bw)
}
