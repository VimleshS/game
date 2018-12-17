package types

import (
	"math/rand"
	"time"
)

type Bowler struct {
	name        string
	intSeeder   int64
	ballsBowled int
	totalRuns   int
}

//NewBolwer ...
func NewBolwer(name string) *Bowler {
	return &Bowler{
		intSeeder: time.Now().UTC().UnixNano(),
		name:      name,
	}
}

func (b *Bowler) StartNewSpell() {
	b.ballsBowled = 0
	b.intSeeder = time.Now().UTC().UnixNano()
	rand.Seed(b.intSeeder)
}

func (b *Bowler) PlayerInfo() string {
	return b.name
}

func (b *Bowler) BowlVaryingSpeed() int {
	b.ballsBowled++
	return rand.Intn(6)
}

func (b *Bowler) BallBowledInSpell() int {
	return b.ballsBowled
}

func (b *Bowler) IsOverComplete() bool {
	return b.ballsBowled == 6
}
