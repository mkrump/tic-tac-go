package uis

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/sc2nomore/tic-tac-go/core/games"
	"github.com/sc2nomore/tic-tac-go/mocks"
)

func TestRenderingEmpty3x3Board(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("O")
	mockPlayer2 := &mocks.Player{}
	mockPlayer2.On("Symbol").Return("X")
	mockPlayers := games.MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Board{}
	mockBoard.On("BoardState").Return([]int{
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
	})
	mockBoard.On("GridSize").Return(3)

	styler := SimpleStyle{}
	boardRender := MakeTTTBoardRender(styler)

	expected :=
		"   0   |   1   |   2   \n" +
			"- - - - - - - - - - - -\n" +
			"   3   |   4   |   5   \n" +
			"- - - - - - - - - - - -\n" +
			"   6   |   7   |   8   \n"

	assert.Equal(t, expected, boardRender.RenderBoard(mockBoard, mockPlayers), "")
}

func TestRenderingNonEmpty3x3Board(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("X")
	mockPlayer2 := &mocks.Player{}
	mockPlayer2.On("Symbol").Return("O")
	mockPlayers := games.MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Board{}
	mockBoard.On("BoardState").Return([]int{
		0, 1, 0,
		3, -1, 5,
		6, 7, 1,
	})
	mockBoard.On("GridSize").Return(3)

	styler := SimpleStyle{}
	boardRender := MakeTTTBoardRender(styler)

	expected :=
		"   0   |   O   |   2   \n" +
			"- - - - - - - - - - - -\n" +
			"   3   |   X   |   5   \n" +
			"- - - - - - - - - - - -\n" +
			"   6   |   7   |   O   \n"

	assert.Equal(t, expected, boardRender.RenderBoard(mockBoard, mockPlayers), "")
}

func TestRenderingNonEmpty4x4Board(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("X")
	mockPlayer2 := &mocks.Player{}
	mockPlayer2.On("Symbol").Return("O")
	mockPlayers := games.MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Board{}
	mockBoard.On("BoardState").Return([]int{
		0, 1, 0, 0,
		0, -1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 0,
	})
	mockBoard.On("GridSize").Return(4)

	styler := SimpleStyle{}
	boardRender := MakeTTTBoardRender(styler)

	expected :=
		"   0   |   O   |   2   |   3   \n" +
			"- - - - - - - - - - - - - - - -\n" +
			"   4   |   X   |   6   |   7   \n" +
			"- - - - - - - - - - - - - - - -\n" +
			"   8   |   9   |   O   |  11   \n" +
			"- - - - - - - - - - - - - - - -\n" +
			"  12   |  13   |  14   |  15   \n"

	assert.Equal(t, expected, boardRender.RenderBoard(mockBoard, mockPlayers), "")
}
