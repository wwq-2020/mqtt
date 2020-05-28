package header

import "sync"

var (
	pool = &sync.Pool{
		New: func() interface{} {
			return New()
		},
	}
)

// Get Get
func Get() Header {
	header := pool.Get().(*header)
	header.Reset()
	return header
}

// Put Put
func Put(header Header) {
	pool.Put(header)
}
