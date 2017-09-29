package main

import (
	"fmt"
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
	console_ui.RenderBoard()
	for {
		console.RequestUserMove(os.Stdout)
		err := console_ui.GetMove()
		if err != nil {
			continue
		}
		console_ui.RenderBoard()
		if game.IsWin() {
			console.ConsolePrint(fmt.Sprintf("%s wins!", game.InActivePlayerMarker()))
			break
		}
		if game.IsTie() {
			console.ConsolePrint(fmt.Sprintf("Tie..."))
			break
		}
	}
}
