package uis

import (
	"github.com/sc2nomore/tic-tac-go/uis"
	"strconv"
)

func ValidateMove(userMove interface{}) (int, error) {
	switch move := userMove.(type) {
	case int:
		return move, nil
	case string:
		return convertToInt(move)
	default:
		return 0, uis.ErrInvalidInput
	}
}

func convertToInt(input string) (int, error) {
	move, err := strconv.Atoi(input)
	if err != nil {
		return 0, uis.ErrInvalidInput
	}
	return move - 1, err
}
