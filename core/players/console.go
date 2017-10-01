package players

import (
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/consolettt"
)

type ConsoleStrategy struct {
	console consolettt.Console
}

func MakeConsoleStrategy(console consolettt.Console) ConsoleStrategy {
	return ConsoleStrategy{
		console: console,
	}
}

func (consoleStrategy ConsoleStrategy) FindMove(core.Board, int) interface{} {
	consoleStrategy.console.RenderMessage("Select an open square: ")
	return consoleStrategy.console.ReadInput()
}
