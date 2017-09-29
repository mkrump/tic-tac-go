package core

type Game interface {
	GameBoard() Board
	GamePlayers() PlayerMapper
	GetMove() interface{}
	MakeMove(move int) error
}
