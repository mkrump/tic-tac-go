package core

type Player interface {
	Symbol() string
	Move(board Board, boardActive int) interface{}
}


