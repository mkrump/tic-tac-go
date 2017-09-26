package playertypes

import (
	"math/rand"
)

type ComputerPlayer struct {
	symbol string
}

func (computerPlayer ComputerPlayer) Symbol() string {
	return computerPlayer.symbol
}

func (computerPlayer ComputerPlayer) Move() interface{} {
	return rand.Intn(9)
}

func MakeComputerPlayer(symbol string) ComputerPlayer {
	return ComputerPlayer{
		symbol: symbol,
	}
}
