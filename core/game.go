package core

type Game interface {
	GameBoard() Board
	GamePlayers() PlayerMapper
	GetMove() interface{}
	MakeMove(move int) error
	IsWin() bool
	IsTie() bool
	InActivePlayerMarker() string
	ActivePlayerMarker() string
}
