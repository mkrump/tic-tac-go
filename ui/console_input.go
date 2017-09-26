package ui

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

//InvalidInputError is error type raised when console input not allowed
type InvalidInputError struct {
	input string
}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("%s is not valid input", string(e.input))
}

//GetUserMove gets user move from io.Reader and converts to int an error is raised
//if input can not be converted to an int
func GetUserMove(in io.Reader) (int, error) {
	re := regexp.MustCompile("\r?\n")
	reader := bufio.NewReader(in)
	input := readInput(reader, re)
	move, err := convertInput(input)
	if err != nil {
		return 0, err
	}
	return move, nil
}

//RequestUserMove writes to io.Writer message asking user for move
func RequestUserMove(out io.Writer) {
	io.WriteString(out, "Select an open square: ")
}

func convertInput(input string) (int, error) {
	move, err := strconv.Atoi(input)
	return move - 1, err
}

func readInput(reader *bufio.Reader, re *regexp.Regexp) string {
	str, _ := reader.ReadString('\n')
	input := re.ReplaceAllString(str, "")
	return input
}
