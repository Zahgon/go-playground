package syncx

import "sync"

type void = struct{}

type Set[T comparable] struct {
	m *sync.Map
}

func NewSet[T comparable]() Set[T] { _ = "STUB: not implemented"; return nil }

func (s Set[T]) Delete(item T) { _ = "STUB: not implemented"; return }

func (s Set[T]) Has(item T) bool { _ = "STUB: not implemented"; return false }

func (s Set[T]) Add(item T) { _ = "STUB: not implemented"; return }

type Map[K comparable, V any] struct {
	m *sync.Map
}

func NewMap[K comparable, V any]() Map[K, V] { _ = "STUB: not implemented"; return nil }

func (m Map[K, V]) Keys() []K { _ = "STUB: not implemented"; return nil }

func (m Map[K, V]) Put(key K, val V) { _ = "STUB: not implemented"; return }

func (m Map[K, V]) Delete(item K) { _ = "STUB: not implemented"; return }

func (m Map[K, V]) Has(item K) bool { _ = "STUB: not implemented"; return false }

func (m Map[K, V]) Get(item K) (res V, ok bool) { _ = "STUB: not implemented"; return *new(V), false }
