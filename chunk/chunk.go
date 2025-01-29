package chunk

import (
	"fmt"

	"github.com/downflux/go-flowfield/vector"
)

var (
	Size = vector.V{8, 8}
	b    = vector.B{vector.V{0, 0}, Size}
)

type O[T any] struct {
	Dimension vector.V

	Wall T
}

type C[T any] struct {
	dimension vector.V

	values [][]T // (x, y)
}

func New[T any](o O[T]) (*C[T], error) {
	if !b.In(o.Dimension) {
		return nil, fmt.Errorf("invalid chunk dimension specified: %v", o.Dimension)
	}

	vs := [][]T{}
	for i := 0; i < Size.X(); i++ {
		vs = append(vs, make([]T, Size.Y()))
	}

	for i := o.Dimension.X(); i < Size.X(); i++ {
		for j := o.Dimension.Y(); j < Size.Y(); j++ {
			vs[i][j] = o.Wall
		}
	}

	return &C[T]{
		dimension: Size,

		values: vs,
	}, nil
}

func (c *C[T]) V(p vector.V) (T, error) {
	if !b.In(p) {
		var empty T
		return empty, fmt.Errorf("invalid relative chunk position: %v", p)
	}
	return c.values[p.X()][p.Y()], nil
}
