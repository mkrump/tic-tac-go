package game

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConsoleInputMac(t *testing.T) {
	inBuffer := bytes.NewBufferString("1\n")
	move, _ := GetUserMove(inBuffer)
	assert.Equal(t, 1, move)
}

func TestGetConsoleInvalidInput(t *testing.T) {
	inBuffer := bytes.NewBufferString("A\r\n")
	_, err := GetUserMove(inBuffer)
	assert.NotNil(t, err)
}

func TestGetConsoleInputWindows(t *testing.T) {
	inBuffer := bytes.NewBufferString("2\r\n")
	move, _ := GetUserMove(inBuffer)
	assert.Equal(t, 2, move)
}

func TestRequestConsoleInput(t *testing.T) {
	outBuffer := bytes.NewBufferString("")
	RequestUserMove(outBuffer)
	assert.Equal(t, "Select an open square: ", outBuffer.String())
}
