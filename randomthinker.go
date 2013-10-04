package main

import (
	"github.com/couchbaselabs/logg"
	cbot "github.com/tleyden/checkers-bot"
)

type RandomThinker struct {
	ourTeamId int
}

func (r RandomThinker) Think(gameState cbot.GameState) (bestMove cbot.ValidMove) {
	ourTeam := gameState.Teams[r.ourTeamId]
	allValidMoves := ourTeam.AllValidMoves()
	randomValidMoveIndex := cbot.RandomIntInRange(0, len(allValidMoves))
	bestMove = allValidMoves[randomValidMoveIndex]
	return
}

func init() {
	logg.LogKeys["MAIN"] = true
}

func main() {
	thinker := &RandomThinker{}
	thinker.ourTeamId = cbot.RED_TEAM
	game := cbot.NewGame(cbot.RED_TEAM, thinker)
	game.GameLoop()
}
