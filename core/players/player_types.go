package players

import (
	"os"
	"github.com/sc2nomore/tic-tac-go/core/tictactoe"
	"github.com/sc2nomore/tic-tac-go/core"
)

func MakeConsolePlayer(symbol string) core.Player {
	return MakeTTTPlayer(
		symbol,
		MakeConsoleStrategy(os.Stdin),
	)
}

func MakeComputerPlayer(symbol string) core.Player {
	return MakeTTTPlayer(
		symbol,
		NegaMaxStrategyAB{Rules: tictactoe.TTTRules{}},
	)
}
