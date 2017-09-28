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

	expected :=
		"   1   |   2   |   3   \n" +
			"- - - - - - - - - - - -\n" +
			"   4   |   5   |   6   \n" +
			"- - - - - - - - - - - -\n" +
			"   7   |   8   |   9   \n"

	assert.Equal(t, expected, RenderBoard(board, mockPlayers), "")
}

func TestRenderingNonEmpty3x3Board(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("X")
	mockPlayer2 := &mocks.Player{}
	mockPlayer2.On("Symbol").Return("O")
	mockPlayers := core.MakePlayers(mockPlayer1, mockPlayer2)
	board := boards.MakeTTTBoard(3)

	board.MakeMove(1, 1)
	board.MakeMove(4, -1)
	board.MakeMove(8, 1)

	expected :=
		"   1   |   O   |   3   \n" +
			"- - - - - - - - - - - -\n" +
			"   4   |   X   |   6   \n" +
			"- - - - - - - - - - - -\n" +
			"   7   |   8   |   O   \n"
	assert.Equal(t, expected, RenderBoard(board, mockPlayers), "")
}

func TestRenderingNonEmpty4x4Board(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("X")
	mockPlayer2 := &mocks.Player{}
	mockPlayer2.On("Symbol").Return("O")
	mockPlayers := core.MakePlayers(mockPlayer1, mockPlayer2)
	board := boards.MakeTTTBoard(4)

	board.MakeMove(1, 1)
	board.MakeMove(5, -1)
	board.MakeMove(10, 1)

	expected :=
		"   1   |   O   |   3   |   4   \n" +
			"- - - - - - - - - - - - - - - -\n" +
			"   5   |   X   |   7   |   8   \n" +
			"- - - - - - - - - - - - - - - -\n" +
			"   9   |  10   |   O   |  12   \n" +
			"- - - - - - - - - - - - - - - -\n" +
			"  13   |  14   |  15   |  16   \n"

	assert.Equal(t, expected, RenderBoard(board, mockPlayers), "")
}
