package chunk

import (
	"fmt"

	"github.com/downflux/go-flowfield/vector"
)

var (
	size = vector.V{8, 8}
	box  = vector.B{vector.V{0, 0}, size}
)

type O[T any] struct {
	Dimension vector.V

	Fill T
}

type C[T any] struct {
	dimension vector.V

	values [][]T // (x, y)
}

func New[T any](o O[T]) (*C[T], error) {
	if !box.In(o.Dimension) {
		return nil, fmt.Errorf("invalid chunk dimension specified: %v", o.Dimension)
	}

	vs := [][]T{}
	for i := 0; i < size.X(); i++ {
		vs = append(vs, make([]T, size.Y()))
	}

	for i := o.Dimension.X(); i < size.X(); i++ {
		for j := o.Dimension.Y(); j < size.Y(); j++ {
			vs[i][j] = o.Fill
		}
	}

	return &C[T]{
		dimension: size,

		values: vs,
	}, nil
}

func (c *C[T]) V(p vector.V) (T, error) {
	if !box.In(p) {
		var empty T
		return empty, fmt.Errorf("invalid relative chunk position: %v", p)
	}
	return c.values[p.X()][p.Y()], nil
}
