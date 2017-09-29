package core

type PlayerMapper interface {
	Player(player int) Player
	PlayerSymbol(player int) (string, error)
}


