package bytespool

import (
	"bytes"
	"sync"
)

var (
	bytes1KPool = &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 1024))
		},
	}
	bytes4KPool = &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 4*1024))
		},
	}
	bytes1MPool = &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 1024*1024))
		},
	}
)

// Get Get
func Get(size int) *bytes.Buffer {
	if size < 1024 {
		br := bytes1KPool.Get().(*bytes.Buffer)
		br.Reset()
		return br
	}
	if size < 4*1024 {
		br := bytes4KPool.Get().(*bytes.Buffer)
		br.Reset()
		return br
	}
	br := bytes1MPool.Get().(*bytes.Buffer)
	br.Reset()
	return br

}
