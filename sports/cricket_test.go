package sports

import (
	"testing"

	"github.com/VimleshS/game/types"
)

func TestBowling(t *testing.T) {
	bowlingTeam := NewTeam("Chennai", nil, types.NewBolwer("Kapil Dev"))
	sport := NewCricket(nil, bowlingTeam, nil, 4, 40)
	done := make(chan bool)
	ballsBowled := sport.StartBowlingSpell(done)
	event := <-ballsBowled
	if event.BallBowledInSpell != 1 {
		t.Error("Ball sequence is inappropriate, expected 1, got ",
			event.BallBowledInSpell)
	}
	<-ballsBowled
	<-ballsBowled
	<-ballsBowled
	<-ballsBowled
	event = <-ballsBowled
	if event.BallBowledInSpell != 6 {
		t.Error("Ball sequence is inappropriate, expected 1, got ",
			event.BallBowledInSpell)
	}

	done <- true

}
