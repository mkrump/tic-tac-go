package core

import "fmt"

//SquareOccupiedError when move attempted for already occupied square
type SquareOccupiedError struct {
	square int
}

func (e *SquareOccupiedError) Error() string {
	return fmt.Sprintf("square %s is occupied", string(e.square))
}

//Board primitive data type to used for ttt game
type Board struct {
	Size       int
	BoardState []int
}

//MakeBoard makes a new board of total length = gridSize x gridSize
func MakeBoard(size int) Board {
	//
	return Board{
		Size:       size,
		BoardState: make([]int, size*size),
	}
}

func (board *Board) MakeMove(square int, player int) error {
	if board.squareOccupied(square) {
		board.BoardState[square] = player
		return nil
	}
	return &SquareOccupiedError{square: square}
}

func (board Board) squareOccupied(square int) bool {
	return board.BoardState[square] == 0
}
