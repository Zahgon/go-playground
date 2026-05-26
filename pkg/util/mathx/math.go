package mathx

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Max[T Number](a, b T) T { _ = "STUB: not implemented"; return *new(T) }

func Min[T Number](a, b T) T { _ = "STUB: not implemented"; return *new(T) }
