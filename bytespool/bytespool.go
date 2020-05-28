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
	if size <= 1024 {
		br := bytes1KPool.Get().(*bytes.Buffer)
		br.Reset()
		return br
	}
	if size <= 4*1024 {
		br := bytes4KPool.Get().(*bytes.Buffer)
		br.Reset()
		return br
	}
	br := bytes1MPool.Get().(*bytes.Buffer)
	br.Reset()
	return br
}

// Put Put
func Put(br *bytes.Buffer) {
	size := br.Cap()
	if size <= 1024 {
		bytes1KPool.Put(br)
		return
	}
	if size <= 4*1024 {
		bytes4KPool.Put(br)
		return
	}
	bytes1MPool.Put(br)
	return
}
