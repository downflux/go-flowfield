package chunk

import (
	"fmt"

	"github.com/downflux/go-flowfield/vector"
)

type ID int

type O struct {
	Dimension vector.V
	Offset    vector.V
}

type C struct {
	id        ID
	dimension vector.V
	offset    vector.V
	dirty     bool

	weights []uint // (x, y)

	neighbors []*C
}

func New(o O) *C {
	ws := make([]uint, o.Dimension.X()*o.Dimension.Y())

	return &C{
		id:        ID(o.Offset.X()/o.Dimension.X() + o.Offset.Y()),
		dimension: o.Dimension,
		offset:    o.Offset,
		dirty:     true,

		weights: ws,
	}
}

func (c *C) ID() ID { return c.ID() }
func (c *C) local(global vector.V) int {
	buf := vector.V{}
	vector.Sub(global, c.offset, buf)
	return buf.X()/c.dimension.X() + buf.Y()
}

func (c *C) SetNeighbors(ns []*C) { c.neighbors = ns }

func (c *C) SetWeight(global vector.V, w uint) error {
	id := c.local(global)
	if id < 0 || id > len(c.weights) {
		return fmt.Errorf("position p %v not in chunk with offset %v and dimension %v", global, c.offset, c.dimension)
	}

	c.dirty = c.dirty && (c.weights[id] == w)
	c.weights[id] = w
	return nil
}

func (c *C) Weight(global vector.V) (uint, error) {
	id := c.local(global)
	if id < 0 || id > len(c.weights) {
		return 0, fmt.Errorf("position p %v not in chunk with offset %v and dimension %v", global, c.offset, c.dimension)
	}
	return c.weights[id], nil
}
