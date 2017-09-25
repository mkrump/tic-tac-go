package game

import (
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderingEmpty3x3Board(t *testing.T) {
	expected :=
		"   1   |   2   |   3   \n" +
			"- - - - - - - - - - - -\n" +
			"   4   |   5   |   6   \n" +
			"- - - - - - - - - - - -\n" +
			"   7   |   8   |   9   \n"
	board := core.MakeBoard(3)
	assert.Equal(t, expected, RenderBoard(board), "")
}

func TestRenderingNonEmpty3x3Board(t *testing.T) {
	expected :=
		"   1   |   X   |   3   \n" +
			"- - - - - - - - - - - -\n" +
			"   4   |   O   |   6   \n" +
			"- - - - - - - - - - - -\n" +
			"   7   |   8   |   X   \n"
	board := core.MakeBoard(3)
	board.MakeMove(1, 1)
	board.MakeMove(4, -1)
	board.MakeMove(8, 1)
	assert.Equal(t, expected, RenderBoard(board), "")
}

func TestRenderingNonEmpty4x4Board(t *testing.T) {
	expected :=
		"   1   |   X   |   3   |   4   \n" +
			"- - - - - - - - - - - - - - - -\n" +
			"   5   |   O   |   7   |   8   \n" +
			"- - - - - - - - - - - - - - - -\n" +
			"   9   |  10   |   X   |  12   \n" +
			"- - - - - - - - - - - - - - - -\n" +
			"  13   |  14   |  15   |  16   \n"
	board := core.MakeBoard(4)
	board.MakeMove(1, 1)
	board.MakeMove(5, -1)
	board.MakeMove(10, 1)
	assert.Equal(t, expected, RenderBoard(board), "")
}
