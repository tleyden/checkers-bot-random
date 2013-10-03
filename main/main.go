package main

import (
	"github.com/couchbaselabs/logg"
	cbot "github.com/tleyden/checkers-bot"
)

func main() {

	logg.LogKeys["MAIN"] = true

	redTeam := cbot.RED_TEAM
	game := cbot.NewGame(redTeam)
	game.GameLoop()

}
