package consolettt

type Console interface {
	RenderMessage(str string)
	ClearConsole()
	ReadInput() string
}
