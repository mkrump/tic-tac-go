package main

import (
	"github.com/sc2nomore/tic-tac-go/core/games"
	"github.com/sc2nomore/tic-tac-go/core/players"
	"github.com/sc2nomore/tic-tac-go/core/tictactoe"
	"github.com/sc2nomore/tic-tac-go/uis"
	"github.com/sc2nomore/tic-tac-go/uis/console"
	"os"
)

func setup() uis.UI {
	consolePlayer := players.MakeConsolePlayer("X")
	computerPlayer := players.MakeComputerPlayer("O")
	players := games.MakePlayers(computerPlayer, consolePlayer)
	game := games.MakeGame(tictactoe.MakeTTTBoard(3), players, tictactoe.TTTRules{})
	styler := console.ColorStyler{}
	boardRender := console.MakeTTTBoardRender(styler)
	return console.MakeConsoleUI(game, boardRender, os.Stdout)
}

func main() {
	consoleUI := setup()
	playing := true
	var message string
	consoleUI.RenderBoard()
	for playing {
		//TODO inject os.Stdout into consoleUI
		console.RequestUserMove(os.Stdout)
		err := consoleUI.GetMove()
		if err != nil {
			continue
		}
		consoleUI.RenderBoard()
		message, playing = consoleUI.NextGameState()
		consoleUI.RenderMessage(message)
	}
}
