package main

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/core/boards"
	"github.com/sc2nomore/tic-tac-go/core/playertypes"
	"github.com/sc2nomore/tic-tac-go/core/rules"
	"github.com/sc2nomore/tic-tac-go/ui"
	"os"
)

//type PlayerMap struct {
//	PlayerMap map[int]core.Player
//}

func main() {

	//Setup
	consolePlayer := playertypes.MakeConsolePlayer("X")
	computerPlayer := playertypes.MakeComputerPlayer("O")
	players := core.MakePlayers(consolePlayer, computerPlayer)
	rules := rules.TTTRules{}
	game := core.MakeGame(boards.MakeBoard(3), players, rules)

	//Main
	ui.ConsolePrint("\n" + ui.RenderBoard(game.Board, game.Players) + "\n")
	turn := 1
	errorCount := 0
	for {
		println(fmt.Sprintf("errors: %d", errorCount))
		if errorCount > 10 {
			break
		}
		ui.RequestUserMove(os.Stdout)
		move, err := ui.GetUserMove(game.ActivePlayer())
		if err != nil {
			println("ERROR")
			errorCount++
			continue
		}
		if err := game.MakeMove(move); err != nil {
			println("ERROR")
			errorCount++
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
		fmt.Println("turn: ", turn)
		fmt.Println(game.Board.BoardState())
		turn++
	}
}
