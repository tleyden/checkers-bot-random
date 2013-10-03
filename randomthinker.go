package checkers_bot_random

import (
	cbot "github.com/tleyden/checkers-bot"
)

type RandomThinker struct {
	ourTeamId int
}

func (r RandomThinker) Think(gameState cbot.GameState) cbot.ValidMove {
	ourTeam := gameState.Teams[r.ourTeamId]
	allValidMoves := ourTeam.AllValidMoves()
	randomValidMoveIndex := cbot.RandomIntInRange(0, len(allValidMoves))
	return allValidMoves[randomValidMoveIndex]
}
