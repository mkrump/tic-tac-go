package core

import (
	"github.com/sc2nomore/tic-tac-go/core/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrentPlayer3x3(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("X")
	mockPlayer2 := &mocks.Player{}
	mockPlayer2.On("Symbol").Return("O")
	mockPlayers := MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("OpenSquares").Return([]int{
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
	})
	game := TTTGame{Board: mockBoard, Players: mockPlayers}

	assert.Equal(t, "X", game.ActivePlayerMarker())
	assert.Equal(t, "O", game.InActivePlayerMarker())
	assert.Equal(t, mockPlayer1, game.ActivePlayer())
	assert.Equal(t, mockPlayer2, game.InActivePlayer())
}

func TestCurrentPlayer4x4Start(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("X")
	mockPlayer2 := &mocks.Player{}
	mockPlayer2.On("Symbol").Return("O")
	mockPlayers := MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(4)
	mockBoard.On("OpenSquares").Return([]int{
		0, 1, 2, 3,
		4, 5, 6, 7,
		7, 8, 9, 10,
		11, 12, 13, 14,
	})
	game := TTTGame{Board: mockBoard, Players: mockPlayers}

	assert.Equal(t, "X", game.ActivePlayerMarker())
	assert.Equal(t, "O", game.InActivePlayerMarker())
	assert.Equal(t, mockPlayer1, game.ActivePlayer())
	assert.Equal(t, mockPlayer2, game.InActivePlayer())
}

func TestCurrentPlayer4x4AfterMoves(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("X")
	mockPlayer2 := &mocks.Player{}
	mockPlayer2.On("Symbol").Return("O")
	mockPlayers := MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(4)
	mockBoard.On("OpenSquares").Return([]int{
		0,
		4, 5, 6, 7,
		7, 8, 9, 10,
		11, 12, 13, 14,
	})
	game := TTTGame{Board: mockBoard, Players: mockPlayers}

	assert.Equal(t, "O", game.ActivePlayerMarker())
	assert.Equal(t, "X", game.InActivePlayerMarker())
	assert.Equal(t, mockPlayer2, game.ActivePlayer())
	assert.Equal(t, mockPlayer1, game.InActivePlayer())
}

func TestCurrentPlayer3x3AfterMoves(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("X")
	mockPlayer2 := &mocks.Player{}
	mockPlayer2.On("Symbol").Return("O")
	mockPlayers := MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("OpenSquares").Return([]int{
		1, 2,
		3, 5,
		6, 7,
	})
	game := TTTGame{Board: mockBoard, Players: mockPlayers}

	assert.Equal(t, "O", game.ActivePlayerMarker())
	assert.Equal(t, "X", game.InActivePlayerMarker())
	assert.Equal(t, mockPlayer2, game.ActivePlayer())
	assert.Equal(t, mockPlayer1, game.InActivePlayer())
}

func TestGameMakeMove(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer2 := &mocks.Player{}
	mockPlayers := MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Playable{}
	game := TTTGame{Board: mockBoard, Players: mockPlayers}
	mockBoard.On("MakeMove", 1, -1).Return(nil)
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("OpenSquares").Return([]int{
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
	})

	game.MakeMove(1)

	mockBoard.AssertCalled(t, "MakeMove", 1, -1)
}

func TestGameIsTieCalled(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer2 := &mocks.Player{}
	mockPlayers := MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Playable{}
	mockRules := &mocks.Rules{}
	game := TTTGame{Board: mockBoard, Players: mockPlayers, Rules: mockRules}
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
	mockBoard := &mocks.Playable{}
	mockRules := &mocks.Rules{}
	game := TTTGame{Board: mockBoard, Players: mockPlayers, Rules: mockRules}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("OpenSquares").Return([]int{
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
	})
	mockBoard.On("BoardState").Return([]int{})

	mockRules.On("IsWin", mockBoard, 1).Return(false)

	game.IsWin()

	mockRules.AssertCalled(t, "IsWin", mockBoard, 1)
}

func TestGameGetMoveRequestsMoveFromActivePlayer(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer2 := &mocks.Player{}
	mockPlayers := MakePlayers(mockPlayer1, mockPlayer2)
	mockBoard := &mocks.Playable{}
	mockRules := &mocks.Rules{}
	game := TTTGame{Board: mockBoard, Players: mockPlayers, Rules: mockRules}
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
