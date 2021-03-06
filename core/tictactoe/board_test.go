package tictactoe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveOpenSquare3x3(t *testing.T) {
	board := MakeTTTBoard(3)

	board.MakeMove(0, 1)

	expected := make([]int, 3*3)
	expected[0] = 1
	assert.Equal(t, 3, board.GridSize(), "")
	assert.Equal(t, expected, board.BoardState(), "")
}

func TestMoveOpenSquare4x4(t *testing.T) {
	board := MakeTTTBoard(4)

	board.MakeMove(9, -1)

	expected := make([]int, 4*4)
	expected[9] = -1
	assert.Equal(t, 4, board.GridSize(), "")
	assert.Equal(t, expected, board.BoardState(), "")
}

func TestMoveOccupiedSquare4x4(t *testing.T) {
	board := MakeTTTBoard(4)
	board.MakeMove(9, -1)

	err := board.MakeMove(9, 1)

	expected := make([]int, 4*4)
	expected[9] = -1
	assert.NotNil(t, err)
	assert.Equal(t, expected, board.BoardState())
}

func TestMoveOutOfBoundsAbove(t *testing.T) {
	board := MakeTTTBoard(3)

	err := board.MakeMove(100, 1)

	assert.NotNil(t, err)
}

func TestMoveOutOfBoundsBelow(t *testing.T) {
	board := MakeTTTBoard(3)

	err := board.MakeMove(-100, 1)

	assert.NotNil(t, err)
}

func TestOpenSquares(t *testing.T) {
	board := MakeTTTBoard(3)

	board.MakeMove(1, 1)
	board.MakeMove(8, -1)

	expected := []int{0, 2, 3, 4, 5, 6, 7}
	assert.Equal(t, expected, board.OpenSquares())

}

func TestUndoMove(t *testing.T) {
	board := MakeTTTBoard(3)

	board.MakeMove(8, -1)
	board.UndoMove(8)

	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	assert.Equal(t, expected, board.OpenSquares())

}

func TestSetBoardState(t *testing.T) {
	board := MakeTTTBoard(3)
	expected := []int{0, 1, 0, 1, 1, 0, 0, 0, 0}
	board.SetBoardState(expected)

	assert.Equal(t, expected, board.BoardState())

}
