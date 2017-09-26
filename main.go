package main

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/ui"
	"os"
)

type Game struct {
	Players map[int]core.Player
}

func main() {
	board := core.MakeBoard(3)
	game := new(Game)
	ui.ConsolePrint("\n" + ui.RenderBoard(board) + "\n")
	player := -1
	consolePlayer := core.MakeConsolePlayer("X")
	computerPlayer := core.MakeComputerPlayer("O")
	game.Players = map[int]core.Player{-1: consolePlayer, 1: computerPlayer}
	turn := 1
	for {
		ui.RequestUserMove(os.Stdout)
		move, err := ui.GetUserMove(game.Players[player])
		if err != nil {
			println("ERROR")
			continue
		}
		if err := board.MakeMove(move, player); err != nil {
			println("ERROR")
			continue
		}
		println(game.Players[player].Symbol())
		ui.ConsolePrint("\n" + ui.RenderBoard(board) + "\n")
		if core.IsWin(board, player) {
			ui.ConsolePrint(fmt.Sprintf("%s wins!", game.Players[player].Symbol()))
			break
		}
		if core.IsTie(board) {
			ui.ConsolePrint(fmt.Sprintf("Tie..."))
			break
		}
		fmt.Println("turn: ", turn)
		fmt.Println(board.BoardState())
		turn++
		player *= -1
	}
}
