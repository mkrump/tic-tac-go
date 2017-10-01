package consoles

type Console interface {
	RenderMessage(str string)
	ClearConsole()
	ReadInput() string
}




