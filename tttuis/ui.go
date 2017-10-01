package tttuis

type UI interface {
	GetMove() error
	RenderBoard()
	NextGameState() (message string, endGame bool)
	RenderMessage(message string)
}
