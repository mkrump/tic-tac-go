package game

import (
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderingEmpty3x3Board(t *testing.T) {
	expected := "\n" +
		"     |     |     \n" +
		"- - - - - - - - -\n" +
		"     |     |     \n" +
		"- - - - - - - - -\n" +
		"     |     |     \n"
	board := core.MakeBoard(3)
	assert.Equal(t, expected, RenderBoard(board), "")
}

func TestRenderingNonEmpty3x3Board(t *testing.T) {
	expected := "\n" +
		"     |  X  |     \n" +
		"- - - - - - - - -\n" +
		"     |  O  |     \n" +
		"- - - - - - - - -\n" +
		"     |     |  X  \n"
	board := core.MakeBoard(3)
	board.MakeMove(1, 1)
	board.MakeMove(4, -1)
	board.MakeMove(8, 1)
	assert.Equal(t, expected, RenderBoard(board), "")
}

func TestRenderingNonEmpty4x4Board(t *testing.T) {
	expected := "\n" +
		"     |  X  |     |     \n" +
		"- - - - - - - - - - - -\n" +
		"     |  O  |     |     \n" +
		"- - - - - - - - - - - -\n" +
		"     |     |  X  |     \n" +
		"- - - - - - - - - - - -\n" +
		"     |     |     |     \n"
	board := core.MakeBoard(4)
	board.MakeMove(1, 1)
	board.MakeMove(5, -1)
	board.MakeMove(10, 1)
	assert.Equal(t, expected, RenderBoard(board), "")
}
