package main

import (
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/game"
	"os"
)

func main() {
	board := core.MakeBoard(3)
	player := 1
	game.ConsolePrint(game.RenderBoard(board))
	for {
		game.RequestUserMove(os.Stdout)
		move, err := game.GetUserMove(os.Stdin)
		if err != nil {
			println("ERROR")
			continue
		}
		board.MakeMove(move, player)
		player *= -1
		game.ConsolePrint(game.RenderBoard(board))
	}
}
