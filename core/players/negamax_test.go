package players

import (
	"github.com/sc2nomore/tic-tac-go/core/tictactoe"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNegamaxCornerImpliesCenter1(t *testing.T) {
	board := tictactoe.MakeTTTBoard(3)
	boardPosition := []int{
		-1, 0, 0,
		0, 0, 0,
		0, 0, 0,
	}
	board.SetBoardState(boardPosition)
	strat := NegaMaxStrategy{Rules: tictactoe.TTTRules{}}

	move := strat.FindMove(board, 1)

	assert.Equal(t, 4, move)
}

func TestNegamaxCornerImpliesCenter2(t *testing.T) {
	board := tictactoe.MakeTTTBoard(3)
	boardPosition := []int{
		0, 0, 1,
		0, 0, 0,
		0, 0, 0,
	}
	board.SetBoardState(boardPosition)
	strat := NegaMaxStrategy{Rules: tictactoe.TTTRules{}}

	move := strat.FindMove(board, -1)

	assert.Equal(t, 4, move)
}

func TestNegamaxCornerImpliesCenter3(t *testing.T) {
	board := tictactoe.MakeTTTBoard(3)
	boardPosition := []int{
		0, 0, 0,
		0, 0, 0,
		0, 0, -1,
	}
	board.SetBoardState(boardPosition)
	strat := NegaMaxStrategy{Rules: tictactoe.TTTRules{}}

	move := strat.FindMove(board, 1)

	assert.Equal(t, 4, move)
}

func TestNegamaxCornerImpliesCenter4(t *testing.T) {
	board := tictactoe.MakeTTTBoard(3)
	boardPosition := []int{
		0, 0, 0,
		0, 0, 0,
		1, 0, 0,
	}
	board.SetBoardState(boardPosition)
	strat := NegaMaxStrategy{Rules: tictactoe.TTTRules{}}

	move := strat.FindMove(board, -1)

	assert.Equal(t, 4, move)
}

func TestWinIfOppenentDoesntTakeCenter1(t *testing.T) {
	board := tictactoe.MakeTTTBoard(3)
	boardPosition := []int{
		-1, 1, 0,
		0, 0, 0,
		0, 0, 0,
	}
	board.SetBoardState(boardPosition)
	strat := NegaMaxStrategy{Rules: tictactoe.TTTRules{}}

	move := strat.FindMove(board, -1)

	assert.Equal(t, 3, move)
}

func TestNegamaxCenterImpliesCorner(t *testing.T) {
	board := tictactoe.MakeTTTBoard(3)
	boardPosition := []int{
		0, 0, 0,
		0, 1, 0,
		0, 0, 0,
	}
	board.SetBoardState(boardPosition)
	strat := NegaMaxStrategy{Rules: tictactoe.TTTRules{}}

	move := strat.FindMove(board, -1)
	expectedMoves := []int{0, 2, 6, 8}

	assert.Contains(t, expectedMoves, move)
}

func TestWinInOne2(t *testing.T) {
	board := tictactoe.MakeTTTBoard(3)
	boardPosition := []int{
		-1, 1, 1,
		1, -1, 0,
		-1, 0, 0,
	}
	board.SetBoardState(boardPosition)
	strat := NegaMaxStrategy{Rules: tictactoe.TTTRules{}}

	move := strat.FindMove(board, 1)

	assert.Equal(t, 8, move)
}

func TestDrawInTwo(t *testing.T) {
	board := tictactoe.MakeTTTBoard(3)
	boardPosition := []int{
		-1, 0, 0,
		-1, 1, 1,
		1, 0, -1,
	}
	board.SetBoardState(boardPosition)
	strat := NegaMaxStrategy{Rules: tictactoe.TTTRules{}}

	move := strat.FindMove(board, -1)

	assert.Equal(t, 2, move)
}

func Test4x4WinInOne(t *testing.T) {
	board := tictactoe.MakeTTTBoard(4)
	boardPosition := []int{
		-1, 1, 0, 0,
		-1, 1, 0, 0,
		-1, 1, 0, 0,
		0, 0, 0, 0,
	}
	board.SetBoardState(boardPosition)
	strat := NegaMaxStrategyAB{Rules: tictactoe.TTTRules{}}

	move := strat.FindMove(board, -1)

	assert.Equal(t, 12, move)
}

// Use `go test` cmd below to to reproduce benchmark results.
// go test -run=XXX -bench=. -benchtime=10s
// Average ns per function call (far right number).
// AlphaBeta is apprx an order of magnitude faster
// BenchmarkNegaMaxTimeToSecondMove-4     	    1000	  16345956 ns/op
// BenchmarkNegaMaxABTimeToSecondMove-4   	   10000	   2042362 ns/op
// BenchmarkNegaMaxTimeToFirstMove-4      	     100	 146498477 ns/op
// BenchmarkNegaMaxABTimeToFirstMove-4    	    2000	  10536147 ns/op
func BenchmarkNegaMaxTimeToSecondMove(b *testing.B) {
	board := tictactoe.MakeTTTBoard(3)
	strat := NegaMaxStrategy{Rules: tictactoe.TTTRules{}}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		board.MakeMove(0, -1)
		strat.FindMove(board, 1)
	}
}

func BenchmarkNegaMaxABTimeToSecondMove(b *testing.B) {
	board := tictactoe.MakeTTTBoard(3)
	strat := NegaMaxStrategyAB{Rules: tictactoe.TTTRules{}}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		board.MakeMove(0, -1)
		strat.FindMove(board, 1)
	}
}

func BenchmarkNegaMaxTimeToFirstMove(b *testing.B) {
	board := tictactoe.MakeTTTBoard(3)
	strat := NegaMaxStrategy{Rules: tictactoe.TTTRules{}}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		strat.FindMove(board, 1)
	}
}

func BenchmarkNegaMaxABTimeToFirstMove(b *testing.B) {
	board := tictactoe.MakeTTTBoard(3)
	strat := NegaMaxStrategyAB{Rules: tictactoe.TTTRules{}}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		strat.FindMove(board, 1)
	}
}
