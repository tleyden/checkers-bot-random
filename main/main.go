package main

import (
	"github.com/couchbaselabs/logg"
	cbot "github.com/tleyden/checkers-bot"
	cbotrandom "github.com/tleyden/checkers-bot-random"
)

func main() {

	logg.LogKeys["MAIN"] = true

	randomThinker := new(cbotrandom.RandomThinker)
	redTeam := cbot.RED_TEAM
	game := cbot.NewGame(redTeam, randomThinker)
	game.GameLoop()

}
