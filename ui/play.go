package ui

import "github.com/sc2nomore/tic-tac-go/core"

// Play starts new tic-tac-toe ui.
func Play() string {
	board := core.MakeBoard(3)
	board.MakeMove(0, 1)
	board.MakeMove(7, -1)
	board.MakeMove(2, -1)
	return RenderBoard(board)
}
