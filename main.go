package main

import (
	"os"

	"github.com/VimleshS/game/score"
	"github.com/VimleshS/game/sports"
	"github.com/VimleshS/game/store"
	"github.com/VimleshS/game/types"
)

func main() {

	pq := store.PlayerQueue{}
	pq.Enqueue(types.NewBatsman("Kirat Bowli", []int{25, 05, 0, 5, 1, 4}))
	pq.Enqueue(types.NewBatsman("R Bumrah", []int{30, 15, 5, 5, 1, 4}))
	pq.Enqueue(types.NewBatsman("N.S Nodhi", []int{40, 20, 5, 10, 1, 4}))
	pq.Enqueue(types.NewBatsman("Shashi Henra", []int{25, 05, 0, 5, 1, 4}))

	sb := scoreboard.Monitor{IoWriter: os.Stdout}
	aggregator := scoreboard.NewAggregate(sb)
	battingTeam := sports.NewTeam("Bangaluru", pq, nil)
	//Reusing the same bowler
	bowlingTeam := sports.NewTeam("Chennai", nil, types.NewBolwer("Kapil Dev"))
	sport := sports.NewCricket(battingTeam, bowlingTeam, aggregator, 4, 40)
	sport.Play()
}
