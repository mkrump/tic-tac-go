package games

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/sc2nomore/tic-tac-go/mocks"
)

func TestPlayerOneCurrentPlayer(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("X")
	mockPlayer2 := &mocks.Player{}
	mockPlayer2.On("Symbol").Return("O")
	mockPlayers := MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Board{}
	mockRules := &mocks.Rules{}
	game := Game{Board: mockBoard, Players: mockPlayers, Rules: mockRules}

	mockRules.On("ActivePlayerNumber", mockBoard).Return(-1)
	mockRules.On("InActivePlayerNumber", mockBoard).Return(1)

	assert.Equal(t, "X", game.ActivePlayerMarker())
	assert.Equal(t, "O", game.InActivePlayerMarker())
	assert.Equal(t, mockPlayer1, game.ActivePlayer())
	assert.Equal(t, mockPlayer2, game.InActivePlayer())
}

func TestPlayerTwoCurrentPlayer(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("X")
	mockPlayer2 := &mocks.Player{}
	mockPlayer2.On("Symbol").Return("O")
	mockPlayers := MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Board{}
	mockRules := &mocks.Rules{}
	game := Game{Board: mockBoard, Players: mockPlayers, Rules: mockRules}

	mockRules.On("ActivePlayerNumber", mockBoard).Return(1)
	mockRules.On("InActivePlayerNumber", mockBoard).Return(-1)

	assert.Equal(t, "O", game.ActivePlayerMarker())
	assert.Equal(t, "X", game.InActivePlayerMarker())
	assert.Equal(t, mockPlayer2, game.ActivePlayer())
	assert.Equal(t, mockPlayer1, game.InActivePlayer())
}

func TestGameMakeMove(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer2 := &mocks.Player{}
	mockPlayers := MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Board{}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("OpenSquares").Return([]int{
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
	})
	mockRules := &mocks.Rules{}
	mockRules.On("ActivePlayerNumber", mockBoard).Return(-1)
	mockRules.On("InActivePlayerNumber", mockBoard).Return(1)
	game := Game{Board: mockBoard, Players: mockPlayers, Rules: mockRules}

	mockBoard.On("MakeMove", 1, -1).Return(nil)
	game.MakeMove(1)

	mockBoard.AssertCalled(t, "MakeMove", 1, -1)
}

func TestGameIsTieCalled(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer2 := &mocks.Player{}
	mockPlayers := MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Board{}
	mockRules := &mocks.Rules{}
	game := Game{Board: mockBoard, Players: mockPlayers, Rules: mockRules}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("OpenSquares").Return([]int{})
	mockBoard.On("BoardState").Return([]int{})

	mockRules.On("IsTie", mockBoard).Return(false)
	game.IsTie()

	mockRules.AssertCalled(t, "IsTie", mockBoard)
}

func TestGameIsWinCalled(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer2 := &mocks.Player{}
	mockPlayers := MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Board{}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("OpenSquares").Return([]int{
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
	})
	mockBoard.On("BoardState").Return([]int{})
	mockRules := &mocks.Rules{}
	mockRules.On("ActivePlayerNumber", mockBoard).Return(-1)
	mockRules.On("InActivePlayerNumber", mockBoard).Return(1)
	game := Game{Board: mockBoard, Players: mockPlayers, Rules: mockRules}

	mockRules.On("IsWin", mockBoard, 1).Return(false)
	game.IsWin()

	mockRules.AssertCalled(t, "IsWin", mockBoard, 1)
}

func TestGameGetMoveRequestsMoveFromActivePlayer(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer2 := &mocks.Player{}
	mockPlayers := MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Board{}
	mockRules := &mocks.Rules{}
	mockRules.On("ActivePlayerNumber", mockBoard).Return(-1)
	mockRules.On("InActivePlayerNumber", mockBoard).Return(1)
	game := Game{Board: mockBoard, Players: mockPlayers, Rules: mockRules}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("OpenSquares").Return([]int{
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
	})

	mockPlayer1.On("Move", mockBoard, -1).Return(1)
	game.GetMove()

	mockPlayer1.AssertCalled(t, "Move", mockBoard, -1)
}
