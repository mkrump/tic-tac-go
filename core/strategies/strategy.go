package strategies

import "github.com/sc2nomore/tic-tac-go/core/boards"

type Strategy interface {
	FindMove(boards.Board, int) interface{}
}
