package types

import (
	"fmt"
	"log"
	"math"
)

type Batsman struct {
	name            string
	ballfaced       int
	totalRuns       int
	wasLastAtCrease bool
	hittingProbs    []int
}

//NewBatsman returns an instance ofBatsman along with probability
func NewBatsman(name string, probs []int) *Batsman {
	return &Batsman{
		name:            name,
		hittingProbs:    probs,
		wasLastAtCrease: true,
	}
}

func (b *Batsman) PlayerInfo() string {
	return b.name
}

func (b *Batsman) Playing(playing bool) {
	b.wasLastAtCrease = playing
}

func (b *Batsman) Run(bowl, srNo int) int {
	weightedInt := float64(bowl) * b.hittingProbability(srNo-1)
	run := int(math.Mod(weightedInt, 7))
	return run
}

func (b *Batsman) hittingProbability(bowl int) float64 {
	if bowl < 0 || bowl > 6 {
		log.Panicln("Probability found negative")
	}
	return float64(b.hittingProbs[bowl])
}

func (b *Batsman) Bowled(run int) bool {
	if run == 0 {
		return true
	}
	return false
}

func (b *Batsman) BallFaced() int {
	return b.ballfaced
}

func (b *Batsman) Print() string {
	star := ""
	if b.wasLastAtCrease {
		star = "*"
	}
	return fmt.Sprintf("%s - %d%s (%d balls)", b.PlayerInfo(), b.ScoredRuns(),
		star, b.BallFaced())
}

func (b *Batsman) ScoredRuns() int {
	return b.totalRuns
}

func (b *Batsman) UpdateStats(run int) {
	b.totalRuns += run
	b.ballfaced++
}
