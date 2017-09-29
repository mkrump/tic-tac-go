package players

import "github.com/sc2nomore/tic-tac-go/core"

type TTTPlayer struct {
	symbol string
	strat  core.Strategy
}

func MakeTTTPlayer(symbol string, strategy core.Strategy) TTTPlayer {
	return TTTPlayer{
		symbol: symbol,
		strat:  strategy,
	}
}

func (player TTTPlayer) Move(board core.Board, boardActive int) interface{} {
	return player.strat.FindMove(board, boardActive)
}

func (player TTTPlayer) Symbol() string {
	return player.symbol
}
