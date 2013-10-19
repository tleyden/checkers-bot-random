package main

import (
	"github.com/couchbaselabs/logg"
	cbot "github.com/tleyden/checkers-bot"
)

type RandomThinker struct {
	ourTeamId cbot.TeamType
}

func (r RandomThinker) Think(gameState cbot.GameState) (bestMove cbot.ValidMove, ok bool) {

	ok = true
	ourTeam := gameState.Teams[r.ourTeamId]
	allValidMoves := ourTeam.AllValidMoves()
	if len(allValidMoves) > 0 {
		randomValidMoveIndex := cbot.RandomIntInRange(0, len(allValidMoves))
		bestMove = allValidMoves[randomValidMoveIndex]
	} else {
		ok = false
	}

	return
}

func (r RandomThinker) GameFinished(gameState cbot.GameState) (shouldQuit bool) {
	shouldQuit = false
	return
}

func init() {
	logg.LogKeys["MAIN"] = true
	logg.LogKeys["DEBUG"] = false
}

func main() {
	team, syncGatewayUrl, feedType := cbot.ParseCmdLine()
	thinker := &RandomThinker{}
	thinker.ourTeamId = team
	game := cbot.NewGame(thinker.ourTeamId, thinker)
	game.SetServerUrl(syncGatewayUrl)
	game.SetFeedType(feedType)
	game.SetDelayBeforeMove(false)
	game.GameLoop()
}
