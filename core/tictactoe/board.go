package tictactoe

import "github.com/sc2nomore/tic-tac-go/core"

//TTTBoard primitive data type to used for ttt uis
type TTTBoard struct {
	size       int
	boardState []int
}

//MakeTTTBoard makes a new TTTBoard of total length = gridSize x gridSize
//have to using MakeTTTBoard
func MakeTTTBoard(size int) TTTBoard {
	//
	return TTTBoard{
		size:       size,
		boardState: make([]int, size*size),
	}
}

func (board *TTTBoard) SetBoardState(boardState []int) {
	board.boardState = boardState
}

//BoardState returns an array of ints representing the game tictactoe
func (board TTTBoard) BoardState() []int {
	return board.boardState
}

//GridSize size of n for nxn tictactoe
func (board TTTBoard) GridSize() int {
	return board.size
}

//MakeMove mutates the TTTBoard object to reflect a player move
//if a square is available else it returns a ErrSquareOccupied
func (board TTTBoard) MakeMove(square int, player int) error {
	switch {
	case board.outOfBounds(square):
		return core.ErrOutOfBounds
	case board.squareOccupied(square):
		return core.ErrSquareOccupied
	default:
		board.boardState[square] = player
		return nil
	}
}

func (board TTTBoard) UndoMove(square int) error {
	if board.outOfBounds(square) {
		return core.ErrOutOfBounds
	}
	board.boardState[square] = 0
	return nil
}

func (board TTTBoard) OpenSquares() []int {
	var openSquares []int
	for i, square := range board.BoardState() {
		if square == 0 {
			openSquares = append(openSquares, i)
		}
	}
	return openSquares
}

func (board TTTBoard) outOfBounds(square int) bool {
	return square < 0 || square >= len(board.BoardState())
}

func (board TTTBoard) squareOccupied(square int) bool {
	return board.boardState[square] != 0
}
