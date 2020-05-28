package codec

import (
	"io"
	"sync"
)

var (
	pool = &sync.Pool{
		New: func() interface{} {
			return New()
		},
	}
)

// Get Get
func Get(rw io.ReadWriter) Codec {
	codec := pool.Get().(*codec)
	codec.Reset(rw)
	return codec
}

// Put Put
func Put(codec Codec) {
	pool.Put(codec)
}
