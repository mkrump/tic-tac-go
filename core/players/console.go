package players

import (
	"github.com/sc2nomore/tic-tac-go/tttuis/consolettt"
	"github.com/sc2nomore/tic-tac-go/core"
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
