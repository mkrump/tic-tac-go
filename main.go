package main

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/core/tictactoe"
	"github.com/sc2nomore/tic-tac-go/core/players"
	"github.com/sc2nomore/tic-tac-go/uis"
	"os"
	"github.com/sc2nomore/tic-tac-go/core/games"
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
	styler := uis.ColorStyler{}
	boardRender := uis.MakeTTTBoardRender(styler)

	console_ui := uis.ConsoleUI{game, boardRender}

	//Main
	console_ui.RenderBoard()
	for {
		uis.RequestUserMove(os.Stdout)
		err := console_ui.GetMove()
		if err != nil {
			continue
		}
		console_ui.RenderBoard()
		if game.IsWin() {
			uis.ConsolePrint(fmt.Sprintf("%s wins!", game.InActivePlayerMarker()))
			break
		}
		if game.IsTie() {
			uis.ConsolePrint(fmt.Sprintf("Tie..."))
			break
		}
	}
}
