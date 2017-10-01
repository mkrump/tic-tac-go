package players

import (
	"github.com/sc2nomore/tic-tac-go/tttuis/consolettt"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/core/tictactoe"
	"os"
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
