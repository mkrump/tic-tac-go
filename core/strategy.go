package core

type Strategy interface {
	FindMove(Board, int) interface{}
}
