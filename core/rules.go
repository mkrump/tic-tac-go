package core

type Rules interface {
	IsWin(board Board, player int) bool
	IsTie(board Board) bool
}