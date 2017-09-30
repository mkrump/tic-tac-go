package menus

import (
	"io"
	"bufio"
	"regexp"
)

type ConsoleMenu struct {
	in  io.Reader
	out io.Writer
}

func NewConsoleMenu(in io.Reader, out io.Writer)  *ConsoleMenu {
	return &ConsoleMenu{
		in: in,
		out: out,
	}
}

func (consoleMenu ConsoleMenu) RenderMessage(str string) {
	io.WriteString(consoleMenu.out, str)
}

func (consoleMenu ConsoleMenu) ReadInput() string {
	reader := bufio.NewReader(consoleMenu.in)
	re := regexp.MustCompile("\r?\n")
	str, _ := reader.ReadString('\n')
	return re.ReplaceAllString(str, "")
}

func (consoleMenu ConsoleMenu) ReadInput() string {
	reader := bufio.NewReader(consoleMenu.in)
	re := regexp.MustCompile("\r?\n")
	str, _ := reader.ReadString('\n')
	return re.ReplaceAllString(str, "")
}