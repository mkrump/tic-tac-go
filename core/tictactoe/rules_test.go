package tictactoe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test3InRow0IsWin(t *testing.T) {
	board := MakeTTTBoard(3)
	board.SetBoardState([]int{
		1, 1, 1,
		0, 0, 0,
		0, 0, 0,
	})
	tttRules := TTTRules{}

	assert.False(t, tttRules.IsTie(board), "")
	assert.True(t, tttRules.IsWin(board, 1), "")
}

func Test3InRow1IsWin(t *testing.T) {
	board := MakeTTTBoard(3)
	board.SetBoardState([]int{
		0, 0, 0,
		1, 1, 1,
		0, 0, 0,
	})
	tttRules := TTTRules{}

	assert.False(t, tttRules.IsTie(board), "")
	assert.True(t, tttRules.IsWin(board, 1), "")
}

func Test3InRow2IsWin(t *testing.T) {
	board := MakeTTTBoard(3)
	board.SetBoardState([]int{
		0, 0, 0,
		0, 0, 0,
		1, 1, 1,
	})
	tttRules := TTTRules{}

	assert.False(t, tttRules.IsTie(board), "")
	assert.True(t, tttRules.IsWin(board, 1), "")
}

func Test3InColumn0IsWin(t *testing.T) {
	board := MakeTTTBoard(3)
	board.SetBoardState([]int{
		1, 0, 0,
		1, 0, 0,
		1, 0, 0,
	})
	tttRules := TTTRules{}

	assert.False(t, tttRules.IsTie(board), "")
	assert.True(t, tttRules.IsWin(board, 1), "")
}

func Test3InColumn1IsWin(t *testing.T) {
	board := MakeTTTBoard(3)
	board.SetBoardState([]int{
		0, 1, 0,
		0, 1, 0,
		0, 1, 0,
	})
	tttRules := TTTRules{}

	assert.False(t, tttRules.IsTie(board), "")
	assert.True(t, tttRules.IsWin(board, 1), "")
}

func Test3InColumn2IsWin(t *testing.T) {
	board := MakeTTTBoard(3)
	board.SetBoardState([]int{
		0, 0, 1,
		0, 0, 1,
		0, 0, 1,
	})
	tttRules := TTTRules{}

	assert.False(t, tttRules.IsTie(board), "")
	assert.True(t, tttRules.IsWin(board, 1), "")
}

func Test3InDownDiagIsWin(t *testing.T) {
	board := MakeTTTBoard(3)
	board.SetBoardState([]int{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	})
	tttRules := TTTRules{}

	assert.False(t, tttRules.IsTie(board), "")
	assert.True(t, tttRules.IsWin(board, 1), "")
}

func Test3InUpDiagIsWin(t *testing.T) {
	board := MakeTTTBoard(3)
	board.SetBoardState([]int{
		0, 0, 1,
		0, 1, 0,
		1, 0, 0,
	})
	tttRules := TTTRules{}

	assert.False(t, tttRules.IsTie(board), "")
	assert.True(t, tttRules.IsWin(board, 1), "")
}

func Test4InUpDiagIsWin(t *testing.T) {
	board := MakeTTTBoard(4)
	board.SetBoardState([]int{
		0, 0, 0, 1,
		0, 0, 1, 0,
		0, 1, 0, 0,
		1, 0, 0, 0,
	})
	tttRules := TTTRules{}

	assert.False(t, tttRules.IsTie(board), "")
	assert.True(t, tttRules.IsWin(board, 1), "")
}

func Test4InDownDiagIsWin(t *testing.T) {
	board := MakeTTTBoard(4)
	board.SetBoardState([]int{
		-1, 0, 0, 0,
		0, -1, 0, 0,
		0, 0, -1, 0,
		0, 0, 0, -1,
	})
	tttRules := TTTRules{}

	assert.False(t, tttRules.IsTie(board), "")
	assert.True(t, tttRules.IsWin(board, -1), "")
}

func Test4InColumn0IsWin(t *testing.T) {
	board := MakeTTTBoard(4)
	board.SetBoardState([]int{
		-1, 0, 0, 0,
		-1, 0, 0, 0,
		-1, 0, 0, 0,
		-1, 0, 0, 0,
	})
	tttRules := TTTRules{}

	assert.False(t, tttRules.IsTie(board), "")
	assert.True(t, tttRules.IsWin(board, -1), "")
}

func Test4InRow1IsWin(t *testing.T) {
	board := MakeTTTBoard(4)
	board.SetBoardState([]int{
		0, 0, 0, 0,
		-1, -1, -1, -1,
		0, 0, 0, 0,
		0, 0, 0, 0,
	})
	tttRules := TTTRules{}

	assert.False(t, tttRules.IsTie(board), "")
	assert.True(t, tttRules.IsWin(board, -1), "")
}

func Test3IsTie(t *testing.T) {
	board := MakeTTTBoard(3)
	board.SetBoardState([]int{
		-1, 1, -1,
		1, -1, -1,
		1, -1, 1,
	})
	tttRules := TTTRules{}

	assert.True(t, tttRules.IsTie(board), "")
}

func Test4IsTie(t *testing.T) {
	board := MakeTTTBoard(4)
	board.SetBoardState([]int{
		-1, 1, -1, 1,
		1, -1, 1, -1,
		-1, 1, -1, 1,
		-1, 1, -1, 1,
	})
	tttRules := TTTRules{}

	assert.True(t, tttRules.IsTie(board), "")
}

func Test1ActivePlayerNumber3x3(t *testing.T) {
	board := MakeTTTBoard(3)
	board.SetBoardState([]int{
		-1, 1, -1,
		0, 0, 0,
		0, 0, 0,
	})
	tttRules := TTTRules{}

	assert.Equal(t, tttRules.ActivePlayerNumber(board), 1)
	assert.Equal(t, tttRules.InActivePlayerNumber(board), -1)
}

func Test2ActivePlayerNumber3x3(t *testing.T) {
	board := MakeTTTBoard(3)
	board.SetBoardState([]int{
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
	})
	tttRules := TTTRules{}

	assert.Equal(t, tttRules.ActivePlayerNumber(board), -1)
	assert.Equal(t, tttRules.InActivePlayerNumber(board), 1)
}

func Test1ActivePlayerNumber4x4(t *testing.T) {
	board := MakeTTTBoard(4)
	board.SetBoardState([]int{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	})
	tttRules := TTTRules{}

	assert.Equal(t, tttRules.ActivePlayerNumber(board), -1)
	assert.Equal(t, tttRules.InActivePlayerNumber(board), 1)
}

func Test2ActivePlayerNumber4x4(t *testing.T) {
	board := MakeTTTBoard(4)
	board.SetBoardState([]int{
		-1, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	})
	tttRules := TTTRules{}

	assert.Equal(t, tttRules.ActivePlayerNumber(board), 1)
	assert.Equal(t, tttRules.InActivePlayerNumber(board), -1)
}
