package uis

type UI interface {
	GetMove() error
	RenderBoard()
}
