package main

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/core/boards"
	"github.com/sc2nomore/tic-tac-go/core/playertypes"
	"github.com/sc2nomore/tic-tac-go/core/rules"
	"github.com/sc2nomore/tic-tac-go/core/strategies"
	"github.com/sc2nomore/tic-tac-go/uis"
	"os"
)

func main() {

	//Setup
	consoleStrategy := strategies.MakeConsoleStrategy(os.Stdin)
	consolePlayer := playertypes.MakeTTTPlayer(
		"X",
		consoleStrategy,
	)
	computerPlayer := playertypes.MakeTTTPlayer(
		"O",
		strategies.NegaMaxStrategyAB{Rules: rules.TTTRules{}},
	)
	players := core.MakePlayers(computerPlayer, consolePlayer)
	game := core.MakeGame(boards.MakeTTTBoard(3), players, rules.TTTRules{})
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
