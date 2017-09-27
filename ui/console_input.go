package ui

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/core/playertypes"
	"io"
	"strconv"
)

//InvalidInputError is error type raised when console input not allowed
type InvalidInputError struct {
	input string
}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("%s is not valid input", e.input)
}

//GetUserMove gets user move from io.Reader and converts to int an error is raised
//if input can not be converted to an int
func GetUserMove(mover playertypes.Player) (int, error) {
	userMove := mover.Move()
	switch move := userMove.(type) {
	case int:
		return move, nil
	case string:
		return convertInput(move)
	default:
		return 0, &InvalidInputError{}
	}
}

//RequestUserMove writes to io.Writer message asking user for move
func RequestUserMove(out io.Writer) {
	io.WriteString(out, "Select an open square: ")
}

func convertInput(input string) (int, error) {
	move, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return move - 1, err
}
