package main

import (
	"github.com/couchbaselabs/logg"
	cbot "github.com/tleyden/checkers-bot"
	cbotrandom "github.com/tleyden/checkers-bot-random"
)

func init() {
	logg.LogKeys["MAIN"] = true
}

func main() {
	thinker := new(cbotrandom.RandomThinker)
	thinker.ourTeamId = cbot.RED_TEAM
	game := cbot.NewGame(cbot.RED_TEAM, thinker)
	game.GameLoop()
}
