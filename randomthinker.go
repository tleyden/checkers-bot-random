package main

import (
	"flag"
	"fmt"
	"github.com/couchbaselabs/logg"
	cbot "github.com/tleyden/checkers-bot"
	"net/http"
	"net/url"
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

func parseCmdLine() (team int, syncGatewayUrl string, proxyPort *int) {

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
	proxyPort = flag.Int(
		"proxyPort",
		-1,
		"The proxy port if you want to use a proxy",
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

func possiblyConfigureProxy(proxyPort *int) {
	// proxy
	// TODO: I don't think this actually works!!  There is a bunch
	// of traffic I'm not seeing .. and I think the go-couch library
	// is setting up its own transport or something.
	if *proxyPort == -1 {
		logg.LogTo("MAIN", "No proxy port configured, not using proxy")
	} else {
		proxyUrlStr := fmt.Sprintf("http://localhost:%d", *proxyPort)
		logg.LogTo("MAIN", "Connecting via proxy on port: %v", proxyUrlStr)
		proxyUrl, err := url.Parse(proxyUrlStr)
		http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		if err != nil {
			panic("proxy issue")
		}

	}

}

func main() {
	team, syncGatewayUrl, proxyPort := parseCmdLine()
	possiblyConfigureProxy(proxyPort)
	thinker := &RandomThinker{}
	thinker.ourTeamId = team
	game := cbot.NewGame(thinker.ourTeamId, thinker)
	game.SetServerUrl(syncGatewayUrl)
	game.SetDelayBeforeMove(false)
	game.GameLoop()
}
