package main

import (
	"flag"
	"github.com/couchbaselabs/logg"
	cbot "github.com/tleyden/checkers-bot"
)

type RandomThinker struct {
	ourTeamId int
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
}

func parseCmdLine() (team int, syncGatewayUrl string) {

	var teamString = flag.String(
		"team",
		"RED",
		"The team, either 'RED' or 'BLUE'",
	)
	var syncGatewayUrlPtr = flag.String(
		"syncGatewayUrl",
		"http://localhost:4984/checkers",
		"The server URL, eg: http://foo.com:4984/checkers",
	)

	flag.Parse()
	if *teamString == "BLUE" {
		team = cbot.BLUE_TEAM
	} else if *teamString == "RED" {
		team = cbot.RED_TEAM
	} else {
		flag.PrintDefaults()
		panic("Invalid command line args given")
	}

	if syncGatewayUrlPtr == nil {
		flag.PrintDefaults()
		panic("Invalid command line args given")
	}

	syncGatewayUrl = *syncGatewayUrlPtr
	return
}

func main() {
	team, syncGatewayUrl := parseCmdLine()
	thinker := &RandomThinker{}
	thinker.ourTeamId = team
	game := cbot.NewGame(thinker.ourTeamId, thinker)
	game.SetServerUrl(syncGatewayUrl)
	game.SetDelayBeforeMove(false)
	game.GameLoop()
}
