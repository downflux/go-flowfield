package field

import (
	"github.com/downflux/go-flowfield/chunk"
	"github.com/downflux/go-flowfield/vector"
)

type O[T any] struct {
	Box vector.B

	Wall T
}

type M[T any] struct {
	box  vector.B
	wall T

	chunks [][]*chunk.C[T] // (x, y)
}

func New[T any](o O[T]) (*M[T], error) {
	chunks := [][]*chunk.C[T]{}
	for i := 0; i < o.Box.D().X()/(chunk.Size.X()+1); i++ {
		chunks = append(chunks, []*chunk.C[T]{})
		for j := 0; j < o.Box.D().Y()/(chunk.Size.Y()+1); j++ {
			if c, err := chunk.New(chunk.O[T]{
				Dimension: vector.V{
					o.Box.D().X() - i*chunk.Size.X(),
					o.Box.D().Y() - i*chunk.Size.Y(),
				},
				Wall: o.Wall,
			}); err != nil {
				return nil, err
			} else {
				chunks[i] = append(chunks[i], c)
			}
		}
	}

	return &M[T]{
		box:  o.Box,
		wall: o.Wall,

		chunks: chunks,
	}, nil
}

func (m *M[T]) V(p vector.V) T {
	if !m.box.In(p) {
		return m.wall
	}
	q := vector.Sub(p, m.box.P())
	q = vector.V{q.X() / chunk.Size.X(), q.Y() / chunk.Size.Y()}
	if v, err := m.chunks[q.X()][q.Y()].V(q); err != nil {
		panic(err)
	} else {
		return v
	}
}
