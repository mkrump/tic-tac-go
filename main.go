package main

import (
	"github.com/sc2nomore/tic-tac-go/core/tictactoe"
	"github.com/sc2nomore/tic-tac-go/core/players"
	"os"
	"github.com/sc2nomore/tic-tac-go/core/games"
	"github.com/sc2nomore/tic-tac-go/uis/console"
)

func main() {

	//Setup
	consoleStrategy := players.MakeConsoleStrategy(os.Stdin)
	consolePlayer := players.MakeTTTPlayer(
		"X",
		consoleStrategy,
	)
	computerPlayer := players.MakeTTTPlayer(
		"O",
		players.NegaMaxStrategyAB{Rules: tictactoe.TTTRules{}},
	)
	players := games.MakePlayers(computerPlayer, consolePlayer)
	game := games.MakeGame(tictactoe.MakeTTTBoard(3), players, tictactoe.TTTRules{})
	styler := console.ColorStyler{}
	boardRender := console.MakeTTTBoardRender(styler)
	console_ui := console.ConsoleUI{game, boardRender}

	//Main
	playing := true
	console_ui.RenderBoard()
	for playing {
		console.RequestUserMove(os.Stdout)
		err := console_ui.GetMove()
		if err != nil {
			continue
		}
		console_ui.RenderBoard()
		playing = console_ui.RenderNextGameState()
	}
}
