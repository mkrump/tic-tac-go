package ui

import (
	"strconv"
	"github.com/sc2nomore/tic-tac-go/tttuis"
)

func ValidateMove(userMove interface{}) (int, error) {
	switch move := userMove.(type) {
	case int:
		return move, nil
	case string:
		return convertToInt(move)
	default:
		return 0, tttuis.ErrInvalidInput
	}
}

func convertToInt(input string) (int, error) {
	move, err := strconv.Atoi(input)
	if err != nil {
		return 0, tttuis.ErrInvalidInput
	}
	return move - 1, err
}
