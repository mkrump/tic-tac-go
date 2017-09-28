package boards

import "fmt"

//SquareOccupiedError when move attempted for already occupied square
type SquareOccupiedError struct {
	square int
}

func (e *SquareOccupiedError) Error() string {
	return fmt.Sprintf("square %s is occupied", string(e.square))
}

//TTTBoard primitive data type to used for ttt ui
type TTTBoard struct {
	size       int
	boardState []int
}

//Board is an interface used to play ttt
//Gridsize size of n for nxn boards
//Boardstate returns an array of ints representing the game boards
//Make move places a place on the boards and returns and error if the move is invalid
type Board interface {
	GridSize() int
	BoardState() []int
	MakeMove(int, int) error
	UndoMove(int) error
	OpenSquares() []int
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

//BoardState returns an array of ints representing the game boards
func (board TTTBoard) BoardState() []int {
	return board.boardState
}

//GridSize size of n for nxn boards
func (board TTTBoard) GridSize() int {
	return board.size
}

//MakeMove mutates the TTTBoard object to reflect a player move
//if a square is available else it returns a SquareOccupiedError
func (board TTTBoard) MakeMove(square int, player int) error {
	if board.inBounds(square) && board.squareOpen(square) {
		board.boardState[square] = player
		return nil
	}
	return &SquareOccupiedError{square: square}
}

func (board TTTBoard) UndoMove(square int) error {
	if board.inBounds(square) {
		board.boardState[square] = 0
	}
	return &SquareOccupiedError{square: square}
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

func (board TTTBoard) inBounds(square int) bool {
	return square >= 0 && square < len(board.BoardState())
}

func (board TTTBoard) squareOpen(square int) bool {
	return board.boardState[square] == 0
}
