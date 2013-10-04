package main

import (
	"github.com/couchbaselabs/logg"
	cbot "github.com/tleyden/checkers-bot"
)

type RandomThinker struct {
	ourTeamId int
}

func (r RandomThinker) Think(gameState cbot.GameState) (bestMove cbot.ValidMove) {
	bestMove = cbot.ValidMove{}
	ourTeam := gameState.Teams[r.ourTeamId]
	allValidMoves := ourTeam.AllValidMoves()
	if len(allValidMoves) > 0 {
		randomValidMoveIndex := cbot.RandomIntInRange(0, len(allValidMoves))
		bestMove = allValidMoves[randomValidMoveIndex]
	}
	return
}

func (r RandomThinker) GameFinished(gameState cbot.GameState) (shouldQuit bool) {
	shouldQuit = false
	return
}

func init() {
	logg.LogKeys["MAIN"] = true
}

func main() {
	thinker := &RandomThinker{}
	thinker.ourTeamId = cbot.BLUE_TEAM
	game := cbot.NewGame(cbot.BLUE_TEAM, thinker)
	game.SetDelayBeforeMove(false)
	game.GameLoop()
}
