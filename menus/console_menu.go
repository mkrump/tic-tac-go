package menus

import (
	"io"
	"regexp"
	"bufio"
	"github.com/sc2nomore/tic-tac-go/uis"
	"strings"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/core/players"
)

type StartupMenu interface {
	PlayerTypePrompt() (string, error)
	PlayerSymbolPrompt() (string, error)
	SelectPlayerType(string, string) core.Player
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

func (consoleMenu ConsoleMenu) PlayerTypePrompt() (string, error) {
	reader := bufio.NewReader(consoleMenu.in)
	choice := consoleMenu.readInput(reader)

	switch choice {
	case "1":
		return "1", nil
	case "2":
		return "2", nil
	default:
		return "0", uis.ErrInvalidOption
	}
}

func (consoleMenu ConsoleMenu) SelectPlayerType(playerType string, playerSymbol string) (core.Player, error) {
	switch playerType {
	case "1":
		return players.MakeConsolePlayer(playerSymbol), nil
	case "2":
		return players.MakeComputerPlayer(playerSymbol), nil
	default:
		return nil, uis.ErrInvalidOption
	}
}

type StartupMenuRunner struct {
	startUpMenu StartupMenu
}

func retry(fn func() (string, error)) string {
	for {
		if val, err := fn(); err == nil {
			return val
		}
	}
}

func (startupMenuRunner StartupMenuRunner) Run() core.Player {
		playerType := retry(startupMenuRunner.startUpMenu.PlayerTypePrompt)
		playerSymbol := retry(startupMenuRunner.startUpMenu.PlayerSymbolPrompt)
		return startupMenuRunner.startUpMenu.SelectPlayerType(playerType, playerSymbol)
}

//
//
//	for{
//	playerType, err := startupMenuRunner.startUpMenu.PlayerTypePrompt()
//	if err == nil{
//	break
//}
//}
//	return startupMenuRunner.startUpMenu.SelectPlayerType(playerType, playerSymbol)
//}
