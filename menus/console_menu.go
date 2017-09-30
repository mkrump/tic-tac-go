package menus

import (
	"io"
	"regexp"
	"bufio"
	"strconv"
	"github.com/sc2nomore/tic-tac-go/uis"
	"strings"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/core/players"
)

func readInput(reader *bufio.Reader) string {
	re := regexp.MustCompile("\r?\n")
	str, _ := reader.ReadString('\n')
	return re.ReplaceAllString(str, "")
}

func PlayerSymbolPrompt(in io.Reader) (string, error) {
	reader := bufio.NewReader(in)
	re := regexp.MustCompile("^[A-z]$")

	input := readInput(reader)
	if re.MatchString(input) {
		return strings.ToUpper(input), nil
	}
	return "", uis.ErrInvalidOption
}

func PlayerTypePrompt(in io.Reader) (int, error) {
	reader := bufio.NewReader(in)
	input := readInput(reader)

	choice, _ := strconv.Atoi(input)
	switch choice {
	case 1:
		return 1, nil
	case 2:
		return 2, nil
	default:
		return 0, uis.ErrInvalidOption
	}
}

func SelectPlayerType(playerType int, playerSymbol string) (core.Player, error) {
	switch playerType {
	case 1:
		return players.MakeConsolePlayer(playerSymbol), nil
	case 2:
		return players.MakeComputerPlayer(playerSymbol), nil
	default:
		return nil, uis.ErrInvalidOption
	}
}
