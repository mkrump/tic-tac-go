package ui

import (
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/core/boards"
	"github.com/sc2nomore/tic-tac-go/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)



func TestRenderingEmpty3x3Board(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("O")
	mockPlayer2 := &mocks.Player{}
	mockPlayer2.On("Symbol").Return("X")
	mockPlayers := core.MakePlayers(mockPlayer1, mockPlayer2)
	board := boards.MakeTTTBoard(3)
	styler := SimpleStyle{}
	expected :=
		"   0   |   1   |   2   \n" +
			"- - - - - - - - - - - -\n" +
			"   3   |   4   |   5   \n" +
			"- - - - - - - - - - - -\n" +
			"   6   |   7   |   8   \n"

	assert.Equal(t, expected, RenderBoard(board, mockPlayers, styler), "")
}

func TestRenderingNonEmpty3x3Board(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("X")
	mockPlayer2 := &mocks.Player{}
	mockPlayer2.On("Symbol").Return("O")
	mockPlayers := core.MakePlayers(mockPlayer1, mockPlayer2)
	board := boards.MakeTTTBoard(3)
	styler := SimpleStyle{}

	board.MakeMove(1, 1)
	board.MakeMove(4, -1)
	board.MakeMove(8, 1)

	expected :=
		"   0   |   O   |   2   \n" +
			"- - - - - - - - - - - -\n" +
			"   3   |   X   |   5   \n" +
			"- - - - - - - - - - - -\n" +
			"   6   |   7   |   O   \n"
	assert.Equal(t, expected, RenderBoard(board, mockPlayers, styler), "")
}

func TestRenderingNonEmpty4x4Board(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("X")
	mockPlayer2 := &mocks.Player{}
	mockPlayer2.On("Symbol").Return("O")
	mockPlayers := core.MakePlayers(mockPlayer1, mockPlayer2)
	board := boards.MakeTTTBoard(4)
	styler := SimpleStyle{}

	board.MakeMove(1, 1)
	board.MakeMove(5, -1)
	board.MakeMove(10, 1)

	expected :=
		"   0   |   O   |   2   |   3   \n" +
			"- - - - - - - - - - - - - - - -\n" +
			"   4   |   X   |   6   |   7   \n" +
			"- - - - - - - - - - - - - - - -\n" +
			"   8   |   9   |   O   |  11   \n" +
			"- - - - - - - - - - - - - - - -\n" +
			"  12   |  13   |  14   |  15   \n"

	assert.Equal(t, expected, RenderBoard(board, mockPlayers, styler), "")
}
