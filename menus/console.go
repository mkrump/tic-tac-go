package menus
//
//import (
//	"io"
//	"bufio"
//	"regexp"
//)
//
//
//type ConsoleMenu struct {
//	in  io.Reader
//	out io.Writer
//}
//
//type Console interface {
//	ReadInput(reader io.Reader)
//	WriteOutput(string)
//}
//
//func (consoleMenu ConsoleMenu) RenderMessage(str string) {
//	io.WriteString(consoleMenu.out, str)
//}
//
//func (consoleMenu ConsoleMenu) readInput(reader io.Reader) string {
//	reader := bufio.NewReader(consoleMenu.in)
//	re := regexp.MustCompile("\r?\n")
//	str, _ := reader.ReadString('\n')
//	return re.ReplaceAllString(str, "")
//}
//
//func (consoleMenu ConsoleMenu)) ReadInput(reader io.Reader)  {
//	return readInput(in)
//}
