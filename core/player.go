package core

import (
	"bufio"
	"io"
	"math/rand"
	"os"
	"regexp"
)

type Player interface {
	Symbol() string
	Move() interface{}
}

type ComputerPlayer struct {
	symbol string
}

func (computerPlayer ComputerPlayer) Symbol() string {
	return computerPlayer.symbol
}

func (computerPlayer ComputerPlayer) Move() interface{} {
	return rand.Intn(9)
}

func MakeComputerPlayer(symbol string) ComputerPlayer {
	return ComputerPlayer{
		symbol: symbol,
	}
}

//----
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
