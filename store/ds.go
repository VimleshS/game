package store

import "github.com/VimleshS/game/types"

type PlayerQueue []*types.Batsman

func (p *PlayerQueue) Dequeue() *types.Batsman {
	_batsman := (*p)[0]
	*p = (*p)[1:]
	return _batsman
}

func (p *PlayerQueue) Enqueue(b *types.Batsman) {
	*p = append(*p, b)
}

func (p *PlayerQueue) HasNext() bool {
	return len(*p) > 0
}

func (p *PlayerQueue) Length() int {
	return len(*p)
}
