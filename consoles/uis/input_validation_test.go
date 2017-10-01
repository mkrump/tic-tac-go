package uis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserMoveInt(t *testing.T) {
	move, _ := ValidateMove(3)

	assert.Equal(t, 3, move)
}

func TestGetUserMoveBadString(t *testing.T) {
	_, err := ValidateMove("A")

	assert.NotNil(t, err)
}

func TestGetUserMoveValidString(t *testing.T) {
	move, _ := ValidateMove("2")

	assert.Equal(t, 1, move)
}

