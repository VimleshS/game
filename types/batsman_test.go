package types

import (
	"reflect"
	"testing"
)

func TestBatsmanForProbability(t *testing.T) {
	actualProb := []int{0, 0, 50, 50, 0, 0}
	batsman := NewBatsman("John D", actualProb)
	if !reflect.DeepEqual(batsman.hittingProbs, actualProb) {
		t.Errorf("probability not matching req %v and found %v", actualProb,
			batsman.hittingProbs)
	}

	bowler := NewBolwer("Johhny D")
	runsScoredForFirstBall := batsman.Run(bowler.BowlVaryingSpeed(), bowler.BallBowledInSpell())
	if runsScoredForFirstBall != 0 {
		t.Errorf("hitting probability is not considered req:  %v and found %v",
			0, runsScoredForFirstBall)
	}

	runsScoredForSecondBall := batsman.Run(bowler.BowlVaryingSpeed(), bowler.BallBowledInSpell())
	if runsScoredForSecondBall != 0 {
		t.Errorf("hitting probability is not considered req:  %v and found %v",
			0, runsScoredForSecondBall)
	}
}

func TestBatsmanForProbabilitySlice(t *testing.T) {
	actualProb := []int{20, 20, 10, 10, 20, 20}
	batsman := NewBatsman("John D", actualProb)

	if batsman.hittingProbs[0] != 20 {
		t.Errorf("probabilities pushed per ball is not considered req:  %v and found %v",
			20, batsman.hittingProbs[0])
	}

}
