package main

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/core/boards"
	"github.com/sc2nomore/tic-tac-go/core/playertypes"
	"github.com/sc2nomore/tic-tac-go/core/rules"
	"github.com/sc2nomore/tic-tac-go/core/strategies"
	"github.com/sc2nomore/tic-tac-go/ui"
	"os"
)

func main() {

	//Setup
	consoleStrategy := strategies.MakeConsoleStrategy(os.Stdin)
	consolePlayer := playertypes.MakeTTTPlayer("X", consoleStrategy)
	computerPlayer := playertypes.MakeTTTPlayer(
		"O",
		strategies.NegaMaxStrategyAB{Rules: rules.TTTRules{}},
	)
	players := core.MakePlayers(consolePlayer, computerPlayer)
	game := core.MakeGame(boards.MakeBoard(3), players, rules.TTTRules{})

	//Main
	ui.ConsolePrint("\n" + ui.RenderBoard(game.Board, game.Players) + "\n")
	for {
		ui.RequestUserMove(os.Stdout)
		move, err := ui.ValidateMove(game.GetMove())
		if err != nil {
			println("ERROR 1")
			continue
		}
		if err := game.MakeMove(move); err != nil {
			println("ERROR 2")
			continue
		}

		ui.ConsolePrint("\n" + ui.RenderBoard(game.Board, game.Players) + "\n")
		if game.IsWin() {
			ui.ConsolePrint(fmt.Sprintf("%s wins!", game.InActivePlayerMarker()))
			break
		}
		if game.IsTie() {
			ui.ConsolePrint(fmt.Sprintf("Tie..."))
			break
		}
	}
}
