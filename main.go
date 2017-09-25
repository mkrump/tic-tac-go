package main

import (
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/game"
	"os"
)

func main() {
	board := core.MakeBoard(3)
	player := 1
	game.ConsolePrint("\n" + game.RenderBoard(board) + "\n")
	for {
		game.RequestUserMove(os.Stdout)
		move, err := game.GetUserMove(os.Stdin)
		if err != nil {
			println("ERROR")
			continue
		}
		if err := board.MakeMove(move, player); err != nil {
			println("ERROR")
			continue
		}
		player *= -1
		game.ConsolePrint("\n" + game.RenderBoard(board) + "\n")
	}
}
