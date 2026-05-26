package buffutil

import (
	"sync"
)

type SlicePool[T any] struct {
	initCap int
	pool    sync.Pool
}

func NewSlicePool[T any](initCap int) *SlicePool[T] { _ = "STUB: not implemented"; return nil }

func (p *SlicePool[T]) Get() []T { _ = "STUB: not implemented"; return nil }

func (p *SlicePool[T]) Put(v []T) {
	_ = "STUB: not implemented"

	//nolint:staticcheck
	return
}
