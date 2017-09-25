package game

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

type InvalidInputError struct {
	input string
}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("%s is not valid input", string(e.input))
}

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

func convertInput(input string) (int, error) {
	move, err := strconv.Atoi(input)
	return move - 1, err
}

func readInput(reader *bufio.Reader, re *regexp.Regexp) string {
	str, _ := reader.ReadString('\n')
	input := re.ReplaceAllString(str, "")
	return input
}

func RequestUserMove(out io.Writer) {
	io.WriteString(out, "Select an open square: ")
}
