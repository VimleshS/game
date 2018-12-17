package sports

import (
	"fmt"

	"github.com/VimleshS/game/score"
	"github.com/VimleshS/game/store"
	"github.com/VimleshS/game/types"
)

type Team struct {
	name          string
	batsmen       store.PlayerQueue
	batsmenPlayed store.PlayerQueue
	playedList    store.PlayerQueue
	bowler        *types.Bowler
	finalScore    []string
	runsScored    int
}

func NewTeam(name string, bq store.PlayerQueue, bowler *types.Bowler) *Team {
	return &Team{name: name, batsmen: bq, bowler: bowler}
}

func (t *Team) Name() string {
	return t.name
}

func (t *Team) UpdateTeamScore(run int) {
	t.runsScored += run
}

func (t *Team) TeamScores() int {
	return t.runsScored
}

func (t *Team) CaptureFinalStats(reqRuns int) {
	if t.runsScored > reqRuns {
		t.finalScore = append(t.finalScore, fmt.Sprintf("%s won by %d wickets",
			t.Name(), len(t.batsmen)+1))
	} else {
		t.finalScore = append(t.finalScore, fmt.Sprintf("%s lost by %d runs",
			t.Name(), reqRuns-t.runsScored))
	}
}

func (t *Team) PrepareSummary(sb scoreboard.Aggregator) {
	for _, p := range t.batsmenPlayed {
		t.finalScore = append(t.finalScore, p.Print())
	}
	sb.Save("", t.finalScore...)
}
