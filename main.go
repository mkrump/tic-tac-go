package main

import (
	"github.com/sc2nomore/tic-tac-go/core/games"
	"github.com/sc2nomore/tic-tac-go/core/tictactoe"
	"github.com/sc2nomore/tic-tac-go/uis"
	"github.com/sc2nomore/tic-tac-go/consoles"
	"os"
	"github.com/sc2nomore/tic-tac-go/consoles/menus"
	uis2 "github.com/sc2nomore/tic-tac-go/consoles/uis"
)

func setup() uis.UI {
	console := consoles.NewTTTConsole(os.Stdin, os.Stdout)
	startupMenu := menus.MakeStartupMenuRunner(menus.NewStartupMenu(console))
	startupMenu.Setup()
	player1, player2 := startupMenu.Players()
	players := games.MakePlayers(player1, player2)
	game := games.MakeGame(tictactoe.MakeTTTBoard(3), players, tictactoe.TTTRules{})
	styler := uis2.ColorStyler{}
	boardRender := uis2.MakeTTTBoardRender(styler)
	return uis2.MakeConsoleUI(game, boardRender, console)
}

func main() {
	consoleUI := setup()
	playing := true
	var message string
	consoleUI.RenderBoard()
	for playing {
		err := consoleUI.GetMove()
		if err != nil {
			continue
		}
		consoleUI.RenderBoard()
		message, playing = consoleUI.NextGameState()
		consoleUI.RenderMessage(message)
	}
}
