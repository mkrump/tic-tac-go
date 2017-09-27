package strategies

import (
	"github.com/sc2nomore/tic-tac-go/core/boards"
	"github.com/sc2nomore/tic-tac-go/core/rules"
	"math"
)

type NegaMaxStrategyAB struct {
	Rules rules.Rules
}

func (strat NegaMaxStrategyAB) negamax(playable boards.Playable, player int,
	depth int, alpha float64, beta float64) float64 {
	if strat.Rules.IsWin(playable, -player) {
		return -1.0 / float64(depth)
	}
	if strat.Rules.IsTie(playable) {
		return 0.0
	}
	bestScore := math.Inf(-1)
	for _, square := range playable.OpenSquares() {
		playable.MakeMove(square, player)
		moveScore := -strat.negamax(playable, -player, depth+1, -beta, -alpha)
		playable.UndoMove(square)
		bestScore = math.Max(moveScore, bestScore)
		alpha = math.Max(moveScore, alpha)
		if alpha >= beta {
			break
		}
	}
	return bestScore
}

func (strat NegaMaxStrategyAB) FindMove(playable boards.Playable, player int) interface{} {
	var bestMove int
	var moveScore float64
	bestScore := math.Inf(-1)
	alpha := math.Inf(-1)
	beta := math.Inf(1)
	for _, square := range playable.OpenSquares() {
		playable.MakeMove(square, player)
		moveScore = -strat.negamax(playable, -player, 1, -beta, -alpha)
		playable.UndoMove(square)
		if moveScore > bestScore {
			bestScore = moveScore
			bestMove = square
		}
	}
	return bestMove
}

type NegaMaxStrategy struct {
	Rules rules.Rules
}

func (strat NegaMaxStrategy) negamax(playable boards.Playable, player int, depth int) float64 {
	if strat.Rules.IsWin(playable, -player) {
		return -1.0 / float64(depth)
	}
	if strat.Rules.IsTie(playable) {
		return 0.0
	}
	bestScore := math.Inf(-1)
	for _, square := range playable.OpenSquares() {
		playable.MakeMove(square, player)
		moveScore := -strat.negamax(playable, -player, depth+1)
		playable.UndoMove(square)
		bestScore = math.Max(moveScore, bestScore)
	}
	return bestScore
}

func (strat NegaMaxStrategy) FindMove(playable boards.Playable, player int) int {
	var bestMove int
	var moveScore float64
	bestScore := math.Inf(-1)
	for _, square := range playable.OpenSquares() {
		playable.MakeMove(square, player)
		moveScore = -strat.negamax(playable, -player, 1)
		playable.UndoMove(square)
		if moveScore > bestScore {
			bestScore = moveScore
			bestMove = square
		}
	}
	return bestMove
}
