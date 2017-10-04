package games

import (
	"github.com/sc2nomore/tic-tac-go/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayerMapPlayerSymbol(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("X")
	playerMap := MakePlayers(mockPlayer1, mockPlayer1)

	player1Symbol, _ := playerMap.PlayerSymbol(1)

	assert.Equal(t, player1Symbol, "X")
}

func TestPlayerMapPlayerSymbolNotFound(t *testing.T) {
	mockPlayer1 := &mocks.Player{}
	mockPlayer1.On("Symbol").Return("X")
	playerMap := MakePlayers(mockPlayer1, mockPlayer1)

	_, err := playerMap.PlayerSymbol(10000)

	assert.NotNil(t, err)
}
