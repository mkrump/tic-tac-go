package core

type Board interface {
	GridSize() int
	BoardState() []int
	MakeMove(int, int) error
	UndoMove(int) error
	OpenSquares() []int
}
