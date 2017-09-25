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
	boardState []int
}

//MakeBoard makes a new board of total length = gridSize x gridSize
func MakeBoard(size int) Board {
	//
	return Board{
		boardState: make([]int, size*size),
	}
}

func (board Board) makeMove(square int, player int) error {
	if board.squareOccupied(square) {
		board.boardState[square] = player
		return nil
	}
	return &SquareOccupiedError{square: square}
}

func (board Board) squareOccupied(square int) bool {
	return board.boardState[square] == 0
}
