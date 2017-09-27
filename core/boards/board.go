package boards

import "fmt"

//SquareOccupiedError when move attempted for already occupied square
type SquareOccupiedError struct {
	square int
}

func (e *SquareOccupiedError) Error() string {
	return fmt.Sprintf("square %s is occupied", string(e.square))
}

//Board primitive data type to used for ttt ui
type Board struct {
	size       int
	boardState []int
}

//Playable is an interface used to play ttt
//Gridsize size of n for nxn boards
//Boardstate returns an array of ints representing the game boards
//Make move places a place on the boards and returns and error if the move is invalid
type Playable interface {
	GridSize() int
	BoardState() []int
	MakeMove(int, int) error
	UndoMove(int) error
	OpenSquares() []int
}

//MakeBoard makes a new Board of total length = gridSize x gridSize
//have to using MakeBoard
func MakeBoard(size int) Board {
	//
	return Board{
		size:       size,
		boardState: make([]int, size*size),
	}
}

func (board *Board) SetBoardState(boardState []int) {
	board.boardState = boardState
}

//BoardState returns an array of ints representing the game boards
func (board Board) BoardState() []int {
	return board.boardState
}

//GridSize size of n for nxn boards
func (board Board) GridSize() int {
	return board.size
}

//MakeMove mutates the Board object to reflect a player move
//if a square is available else it returns a SquareOccupiedError
func (board Board) MakeMove(square int, player int) error {
	if board.inBounds(square) && board.squareOpen(square) {
		board.boardState[square] = player
		return nil
	}
	return &SquareOccupiedError{square: square}
}

func (board Board) UndoMove(square int) error {
	if board.inBounds(square) {
		board.boardState[square] = 0
	}
	return &SquareOccupiedError{square: square}
}

func (board Board) OpenSquares() []int {
	var openSquares []int
	for i, square := range board.BoardState() {
		if square == 0 {
			openSquares = append(openSquares, i)
		}
	}
	return openSquares
}

func (board Board) inBounds(square int) bool {
	return square >= 0 && square < len(board.BoardState())
}

func (board Board) squareOpen(square int) bool {
	return board.boardState[square] == 0
}
