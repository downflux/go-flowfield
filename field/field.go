package field

import (
	"github.com/downflux/go-flowfield/chunk"
	"github.com/downflux/go-flowfield/vector"
)

type O [T]struct {
	Box vector.B

	Fill T
}

type M[T any] struct {
	box vector.B

	chunks [][]C[T]
}

func New[T any](o O) *M[T] {
	vs := [][]T{}
	for i := 0; i < o.Dimension.X(); i++ {
		vs = append(vs, make([]T, o.Dimension.Y()))
	}

	return &M[T]{
		box: o.Box,

		values: vs,
	}
}
