package uis

import (
	"errors"
	"github.com/sc2nomore/tic-tac-go/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestRequestUserMoveValid(t *testing.T) {
	mockGame := &mocks.Game{}
	mockGame.On("GetMove").Return("9")
	mockGame.On("MakeMove", 8).Return(nil)
	consoleUI := MakeConsoleUI(mockGame)

	consoleUI.GetMove()

	mockGame.AssertCalled(t, "GetMove")
}

func TestRequestUserMoveInvalid(t *testing.T) {
	mockGame := &mocks.Game{}
	mockGame.On("GetMove").Return("ABC")
	consoleUI := MakeConsoleUI(mockGame)

	err := consoleUI.GetMove()

	assert.NotNil(t, err)
	mockGame.AssertCalled(t, "GetMove")
}

func TestRequestUserMoveOutBounds(t *testing.T) {
	mockGame := &mocks.Game{}
	mockGame.On("GetMove").Return("1000")
	mockGame.On("MakeMove", mock.AnythingOfType("int")).Return(errors.New("SOME ERROR"))
	consoleUI := MakeConsoleUI(mockGame)

	err := consoleUI.GetMove()

	assert.NotNil(t, err)
	mockGame.AssertCalled(t, "GetMove")
}


