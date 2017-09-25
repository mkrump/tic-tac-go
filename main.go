package main

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/ui"
	"os"
)

type Game struct {
	Players map[int]string
}

func main() {
	board := core.MakeBoard(3)
	game := new(Game)
	game.Players = map[int]string{-1: "X", 1: "O"}
	ui.ConsolePrint("\n" + ui.RenderBoard(board) + "\n")
	player := -1
	for {
		ui.RequestUserMove(os.Stdout)
		move, err := ui.GetUserMove(os.Stdin)
		if err != nil {
			println("ERROR")
			continue
		}
		if err := board.MakeMove(move, player); err != nil {
			println("ERROR")
			continue
		}
		ui.ConsolePrint("\n" + ui.RenderBoard(board) + "\n")
		if core.IsWin(board, player) {
			ui.ConsolePrint(fmt.Sprintf("%s wins!", game.Players[player]))
			break
		}
		if core.IsTie(board) {
			ui.ConsolePrint(fmt.Sprintf("Tie..."))
			break
		}
		player *= -1
	}
}
