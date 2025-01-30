package field

import (
	"fmt"

	"github.com/downflux/go-flowfield/chunk"
	"github.com/downflux/go-flowfield/vector"
)

const (
	wall = 255 // impassible cost field
)

type O struct {
	Dimension vector.V
	ChunkSize vector.V
}

type M struct {
	dimension vector.V
	cost      []*chunk.C
}

func New(o O) (*M, error) {
	if o.Dimension.X()%o.ChunkSize.X() != 0 || o.Dimension.Y()%o.ChunkSize.Y() != 0 {
		return nil, fmt.Errorf("invalid map dimension %v: must be a whole integer multiple of the given chunk size %v", o.Dimension, o.ChunkSize)
	}
	n := vector.V{
		o.Dimension.X() / o.ChunkSize.X(),
		o.Dimension.Y() / o.ChunkSize.Y(),
	}

	l := n.X() * n.Y()
	cost := make([]*chunk.C, l)
	for i := 0; i < n.X(); i++ {
		for j := 0; j < n.Y(); j++ {
			c := chunk.New(chunk.O{
				Dimension: o.ChunkSize,
			})

			cost[i*o.ChunkSize.X()+j] = c
		}
	}

	return &M{
		dimension: o.Dimension,

		cost: cost,
	}, nil
}

func (m *M) V(p vector.V) uint {
	/*
		if !m.box.In(p) {
			return wall
		}
		q := vector.Sub(p, m.box.P())
		q = vector.V{q.X() / chunk.Size.X(), q.Y() / chunk.Size.Y()}
		if v, err := m.chunks[q.X()][q.Y()].V(q); err != nil {
			panic(err)
		} else {
			return v
		}
	*/
	return 0
}
