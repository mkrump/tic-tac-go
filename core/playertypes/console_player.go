package playertypes

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

type ConsolePlayer struct {
	symbol string
	in     io.Reader
}

func (consolePlayer ConsolePlayer) Symbol() string {
	return consolePlayer.symbol
}

func MakeConsolePlayer(symbol string) ConsolePlayer {
	return ConsolePlayer{
		in:     os.Stdin,
		symbol: symbol,
	}
}

func (consolePlayer ConsolePlayer) Move() interface{} {
	re := regexp.MustCompile("\r?\n")
	reader := bufio.NewReader(consolePlayer.in)
	return readInput(reader, re)
}

func readInput(reader *bufio.Reader, re *regexp.Regexp) string {
	str, _ := reader.ReadString('\n')
	input := re.ReplaceAllString(str, "")
	return input
}
