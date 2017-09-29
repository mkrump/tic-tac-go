package players

import (
	"github.com/sc2nomore/tic-tac-go/core"
	"math"
)

type NegaMaxStrategyAB struct {
	Rules core.Rules
}

func (strat NegaMaxStrategyAB) negamax(board core.Board, player int,
	depth int, alpha float64, beta float64) float64 {
	if strat.Rules.IsWin(board, -player) {
		return -1.0 / float64(depth)
	}
	if strat.Rules.IsTie(board) {
		return 0.0
	}
	bestScore := math.Inf(-1)
	for _, square := range board.OpenSquares() {
		board.MakeMove(square, player)
		moveScore := -strat.negamax(board, -player, depth+1, -beta, -alpha)
		board.UndoMove(square)
		bestScore = math.Max(moveScore, bestScore)
		alpha = math.Max(moveScore, alpha)
		if alpha >= beta {
			break
		}
	}
	return bestScore
}

func (strat NegaMaxStrategyAB) FindMove(board core.Board, player int) interface{} {
	var bestMove int
	var moveScore float64
	bestScore := math.Inf(-1)
	alpha := math.Inf(-1)
	beta := math.Inf(1)
	for _, square := range board.OpenSquares() {
		board.MakeMove(square, player)
		moveScore = -strat.negamax(board, -player, 1, -beta, -alpha)
		board.UndoMove(square)
		if moveScore > bestScore {
			bestScore = moveScore
			bestMove = square
		}
	}
	return bestMove
}

type NegaMaxStrategy struct {
	Rules core.Rules
}

// Below is same except w/o Alpha Beta Pruning w/ just leaving in for
// benchmarking purposes. Not used except in if want to rerun benchmarks.
func (strat NegaMaxStrategy) negamax(board core.Board, player int, depth int) float64 {
	if strat.Rules.IsWin(board, -player) {
		return -1.0 / float64(depth)
	}
	if strat.Rules.IsTie(board) {
		return 0.0
	}
	bestScore := math.Inf(-1)
	for _, square := range board.OpenSquares() {
		board.MakeMove(square, player)
		moveScore := -strat.negamax(board, -player, depth+1)
		board.UndoMove(square)
		bestScore = math.Max(moveScore, bestScore)
	}
	return bestScore
}

func (strat NegaMaxStrategy) FindMove(board core.Board, player int) int {
	var bestMove int
	var moveScore float64
	bestScore := math.Inf(-1)
	for _, square := range board.OpenSquares() {
		board.MakeMove(square, player)
		moveScore = -strat.negamax(board, -player, 1)
		board.UndoMove(square)
		if moveScore > bestScore {
			bestScore = moveScore
			bestMove = square
		}
	}
	return bestMove
}
