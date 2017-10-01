package players

import (
	"os"
	"github.com/sc2nomore/tic-tac-go/core/tictactoe"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/consolettt"
)

func MakeConsolePlayer(symbol string) core.Player {
	return MakeTTTPlayer(
		symbol,
		MakeConsoleStrategy(consolettt.NewTTTConsole(os.Stdin, os.Stdout)),
	)
}

func MakeComputerPlayer(symbol string) core.Player {
	return MakeTTTPlayer(
		symbol,
		NegaMaxStrategyAB{Rules: tictactoe.TTTRules{}},
	)
}
