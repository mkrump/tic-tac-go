package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveOpenSquare3x3(t *testing.T) {
	board := MakeBoard(3)
	board.makeMove(0, 1)
	expected := make([]int, 3*3)
	expected[0] = 1
	assert.Equal(t, expected, board.boardState, "")
}

func TestMoveOpenSquare4x4(t *testing.T) {
	board := MakeBoard(4)
	board.makeMove(9, -1)
	expected := make([]int, 4*4)
	expected[9] = -1
	assert.Equal(t, expected, board.boardState, "")
}

func TestMoveOccupiedSquare4x4(t *testing.T) {
	board := MakeBoard(4)
	board.makeMove(9, -1)
	err := board.makeMove(9, 1)
	expected := make([]int, 4*4)
	expected[9] = -1
	assert.NotNil(t, err)
	assert.Equal(t, expected, board.boardState)
}
