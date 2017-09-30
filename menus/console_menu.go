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



type StartupMenu interface {
	PlayerTypePrompt() (int, error)
	PlayerSymbolPrompt(string, error)
	SelectPlayerType(core.Player, error)
}

type ConsoleMenu struct {
	in  io.Reader
	out io.Writer
}

func (consoleMenu ConsoleMenu) RenderMessage(str string) {
	io.WriteString(consoleMenu.out, str)
}

func (consoleMenu ConsoleMenu) readInput(reader *bufio.Reader) string {
	re := regexp.MustCompile("\r?\n")
	str, _ := reader.ReadString('\n')
	return re.ReplaceAllString(str, "")
}

func (consoleMenu ConsoleMenu) PlayerSymbolPrompt() (string, error) {
	reader := bufio.NewReader(consoleMenu.in)
	re := regexp.MustCompile("^[A-z]$")

	input := consoleMenu.readInput(reader)
	if re.MatchString(input) {
		return strings.ToUpper(input), nil
	}
	return "", uis.ErrInvalidOption
}

func (consoleMenu ConsoleMenu) PlayerTypePrompt() (int, error) {
	reader := bufio.NewReader(consoleMenu.in)
	input := consoleMenu.readInput(reader)

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

func (consoleMenu ConsoleMenu) SelectPlayerType(playerType int, playerSymbol string) (core.Player, error) {
	switch playerType {
	case 1:
		return players.MakeConsolePlayer(playerSymbol), nil
	case 2:
		return players.MakeComputerPlayer(playerSymbol), nil
	default:
		return nil, uis.ErrInvalidOption
	}
}

type StartupMenuRunner struct {
	startUpMenu StartupMenu
}

//func (startupMenuRunner StartupMenuRunner) Run(StartupMenu) core.Player {
//	startupMenuRunner.startUpMenu.PlayerTypePrompt()
//
//
//}
