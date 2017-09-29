package players

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/sc2nomore/tic-tac-go/core/mocks"
)

func Test3InRow0IsWin(t *testing.T) {
	strategy := &mocks.Strategy{}
	mockBoard := &mocks.Board{}
	strategy.On("FindMove", mockBoard, 1).Return(true)
	player := MakeTTTPlayer("X", strategy)

	player.Move(mockBoard, 1)

	strategy.AssertCalled(t, "FindMove", mockBoard, 1)
	assert.Equal(t, "X", player.Symbol())
}
