package buffutil

import (
	"bytes"
	"sync"
)

const defaultInitSize = 1024 * 1024

type BufferPool struct {
	buffInitSize int
	pool         sync.Pool
}

func NewBufferPoolWithSize(buffInitSize int) *BufferPool { _ = "STUB: not implemented"; return nil }

func NewBufferPool() *BufferPool { _ = "STUB: not implemented"; return nil }

func (p *BufferPool) Get() RecyclableBuffer {
	_ = "STUB: not implemented"
	return *new(RecyclableBuffer)
}

func (p *BufferPool) Put(b RecyclableBuffer) { _ = "STUB: not implemented"; return }

type RecyclableBuffer struct {
	*bytes.Buffer

	pool *BufferPool
}

func (b RecyclableBuffer) Closed() bool { _ = "STUB: not implemented"; return false }

func (b RecyclableBuffer) Close() error { _ = "STUB: not implemented"; return nil }
