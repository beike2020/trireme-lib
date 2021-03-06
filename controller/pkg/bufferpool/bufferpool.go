package bufferpool

import (
	"sync"
)

// BufferPool implements the interface of httputil.BufferPool in order
// to improve memory utilization in the reverse proxy.
type BufferPool struct {
	s sync.Pool
}

// NewPool creates a new BufferPool.
func NewPool(size int) *BufferPool {
	return &BufferPool{
		s: sync.Pool{
			New: func() interface{} {
				return make([]byte, size)
			},
		},
	}
}

// Get gets a buffer from the pool.
func (b *BufferPool) Get() []byte {
	return b.s.Get().([]byte)
}

// Put returns the buffer to the pool.
func (b *BufferPool) Put(buf []byte) {
	b.s.Put(buf) // nolint
}
