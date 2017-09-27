package playertypes

import (
	"github.com/sc2nomore/tic-tac-go/core/boards"
	"github.com/sc2nomore/tic-tac-go/core/strategies"
)

type Player interface {
	Symbol() string
	Move(board boards.Playable, boardActive int) interface{}
}

type TTTPlayer struct {
	symbol string
	strat  strategies.Strategy
}

func MakeTTTPlayer(symbol string, strategy strategies.Strategy) TTTPlayer {
	return TTTPlayer{
		symbol: symbol,
		strat:  strategy,
	}
}

func (player TTTPlayer) Move(board boards.Playable, boardActive int) interface{} {
	return player.strat.FindMove(board, boardActive)
}

func (player TTTPlayer) Symbol() string {
	return player.symbol
}
