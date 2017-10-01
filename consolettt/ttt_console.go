package consolettt

import (
	"bufio"
	"io"
	"regexp"
)

type TTTConsole struct {
	in  io.Reader
	out io.Writer
}

func (console TTTConsole) ClearConsole() {
	console.RenderMessage("\033c")
}

func NewTTTConsole(in io.Reader, out io.Writer) *TTTConsole {
	return &TTTConsole{
		in:  in,
		out: out,
	}
}

func (console TTTConsole) RenderMessage(str string) {
	io.WriteString(console.out, str)
}

func (console TTTConsole) ReadInput() string {
	reader := bufio.NewReader(console.in)
	re := regexp.MustCompile("\r?\n")
	str, _ := reader.ReadString('\n')
	return re.ReplaceAllString(str, "")
}
