package store

import (
	"testing"

	"github.com/VimleshS/game/types"
)

func TestPlayerQueue(t *testing.T) {
	pq := PlayerQueue{}
	pq.Enqueue(types.NewBatsman("P1", []int{25, 05, 0, 5, 1, 4}))

	if pq.Length() != 1 {
		t.Error("Enqueue method of PlayerQueue failed")
	}

	pq.Enqueue(types.NewBatsman("P2", []int{25, 05, 0, 5, 1, 4}))
	if pq.Length() != 2 {
		t.Error("Enqueue method of PlayerQueue failed")
	}

	e := pq.Dequeue()
	if pq.Length() != 1 {
		t.Error("Dequeue method of PlayerQueue failed")
	}

	if e.PlayerInfo() != "P1" {
		t.Error("Queue implementation is incorrect")
	}

}
