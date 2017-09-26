package ui

import (
	"bytes"
	"github.com/sc2nomore/tic-tac-go/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserMoveInt(t *testing.T) {
	mockPlayer := &mocks.Player{}
	mockPlayer.On("Move").Return(3)
	move, _ := GetUserMove(mockPlayer)
	assert.Equal(t, 3, move)
}

func TestGetUserMoveBadString(t *testing.T) {
	mockPlayer := &mocks.Player{}
	mockPlayer.On("Move").Return("A")
	_, err := GetUserMove(mockPlayer)
	assert.NotNil(t, err)
}

func TestGetUserMoveValidString(t *testing.T) {
	mockPlayer := &mocks.Player{}
	mockPlayer.On("Move").Return("2")
	move, _ := GetUserMove(mockPlayer)
	assert.Equal(t, 1, move)
}

func TestRequestConsoleInput(t *testing.T) {
	outBuffer := bytes.NewBufferString("")
	RequestUserMove(outBuffer)
	assert.Equal(t, "Select an open square: ", outBuffer.String())
}
