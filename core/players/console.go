package players

import (
	"bufio"
	"io"
	"regexp"
	"github.com/sc2nomore/tic-tac-go/core"
)

type ConsoleStrategy struct {
	in io.Reader
}

func MakeConsoleStrategy(in io.Reader) ConsoleStrategy {
	return ConsoleStrategy{
		in: in,
	}
}

func readInput(reader *bufio.Reader, re *regexp.Regexp) string {
	str, _ := reader.ReadString('\n')
	input := re.ReplaceAllString(str, "")
	return input
}

func (consoleStrategy ConsoleStrategy) FindMove(core.Board, int) interface{} {
	re := regexp.MustCompile("\r?\n")
	reader := bufio.NewReader(consoleStrategy.in)
	return readInput(reader, re)
}
