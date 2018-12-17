package sports

import (
	"fmt"

	"github.com/VimleshS/game/score"
	"github.com/VimleshS/game/types"
)

type Cricket struct {
	battingTeam *Team
	bowlingTeam *Team
	striker     *types.Batsman
	nonStriker  *types.Batsman
	aggregator  scoreboard.Aggregator

	runsReqdToWin int
	oversInHand   int
}

func NewCricket(t1, t2 *Team, aggs scoreboard.Aggregator, totalOvers,
	runsToWin int) *Cricket {

	return &Cricket{battingTeam: t1, bowlingTeam: t2,
		aggregator: aggs, oversInHand: totalOvers,
		runsReqdToWin: runsToWin}
}

func (c *Cricket) setGame() {
	//TODO: Validate game layout and fields
	if c.battingTeam.batsmen.Length() < 2 {
		panic("Dont know how play")
	}

	//Set initial players
	c.striker, c.nonStriker =
		c.battingTeam.batsmen.Dequeue(), c.battingTeam.batsmen.Dequeue()
}

func (c *Cricket) checkAndRotateStrike(run, ballCount int) {
	switch run {
	case int(types.One), int(types.Three), int(types.Five):
		c.striker, c.nonStriker = c.nonStriker, c.striker
	}

	switch ballCount {
	case int(types.Six):
		c.striker, c.nonStriker = c.nonStriker, c.striker
	}

}

func (c *Cricket) addPlayerToList(p *types.Batsman) {
	c.battingTeam.batsmenPlayed.Enqueue(p)
}

func (c *Cricket) Play() {
	defer func() {
		c.battingTeam.PrepareSummary(c.aggregator)
		c.aggregator.Dump()
	}()

	runningScore := 0

	c.setGame()
	done := make(chan bool)
	bowlEvents := c.StartBowlingSpell(done)
	for bEvent := range bowlEvents {

		//TODO: Use of typed event may eliminated this added responsibility
		if bEvent.BallBowledInSpell == int(types.One) {
			c.aggregator.Save(fmt.Sprintf("\n%d over left, %d runs to win",
				c.oversInHand-bEvent.RunningOver, c.runsReqdToWin-runningScore))
		}

		score := c.striker.Run(bEvent.Bowl, bEvent.BallBowledInSpell)
		c.striker.UpdateStats(score)
		c.battingTeam.UpdateTeamScore(score)
		runningScore += score

		out := c.striker.Bowled(score)
		if out {
			c.aggregator.Save(fmt.Sprintf("%0.1f %s Bowled out ",
				float64(bEvent.BallBowledInSpell)/10, c.striker.PlayerInfo()))

			c.striker.Playing(false)
			c.addPlayerToList(c.striker)
			if c.battingTeam.batsmen.HasNext() {
				//Replace
				c.striker = c.battingTeam.batsmen.Dequeue()
				c.checkAndRotateStrike(score, bEvent.BallBowledInSpell)
				continue
			} else {
				c.battingTeam.CaptureFinalStats(c.runsReqdToWin)
				c.addPlayerToList(c.nonStriker)
				return
			}
		}

		c.aggregator.Save(fmt.Sprintf("%0.1f %s scores %d runs",
			float64(bEvent.BallBowledInSpell)/10, c.striker.PlayerInfo(), score))

		if runningScore >= c.runsReqdToWin {
			c.doWinningSteps()
			return
		}
		c.checkAndRotateStrike(score, bEvent.BallBowledInSpell)
	}
}

func (c *Cricket) doWinningSteps() {
	c.battingTeam.CaptureFinalStats(c.runsReqdToWin)
	c.addPlayerToList(c.striker)
	c.addPlayerToList(c.nonStriker)
	for c.battingTeam.batsmen.HasNext() {
		_p := c.battingTeam.batsmen.Dequeue()
		_p.Playing(false)
		c.battingTeam.batsmenPlayed.Enqueue(_p)
	}
}

func (c *Cricket) StartBowlingSpell(doneChan chan bool) chan types.Event {
	bowlEventChan := make(chan types.Event)
	go func(doneChan chan bool) {
		defer close(bowlEventChan)
		runningOver := 0

		for runningOver < c.oversInHand {
			c.bowlingTeam.bowler.StartNewSpell()
			for !c.bowlingTeam.bowler.IsOverComplete() {
				ballthrown := c.bowlingTeam.bowler.BowlVaryingSpeed()
				e := types.Event{
					Bowl:              ballthrown,
					BallBowledInSpell: c.bowlingTeam.bowler.BallBowledInSpell(),
					RunningOver:       runningOver,
				}
				select {
				case bowlEventChan <- e:
				case <-doneChan:
					return
				}
			}
			c.bowlingTeam.bowler.StartNewSpell()
			runningOver++
		}
	}(doneChan)
	return bowlEventChan
}
