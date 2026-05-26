package imports

type Queue[T any] struct {
	entries []T
	max     int
}

func NewQueue[T any](size int) *Queue[T] { _ = "STUB: not implemented"; return nil }

func (q *Queue[T]) Occupied() bool { _ = "STUB: not implemented"; return false }

func (q *Queue[T]) MaxOccupancy() int { _ = "STUB: not implemented"; return 0 }

func (q *Queue[T]) Pop() (val T, ok bool) { _ = "STUB: not implemented"; return *new(T), false }

func (q *Queue[T]) Add(items ...T) { _ = "STUB: not implemented"; return }
