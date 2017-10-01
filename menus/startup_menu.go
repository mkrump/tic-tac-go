package menus

import "github.com/sc2nomore/tic-tac-go/core"

type StartupMenu interface {
	PlayerTypePrompt(playerNumber int) (string, error)
	PlayerSymbolPrompt() (string, error)
	SelectPlayerType(string, string) (core.Player, error)
	ClearMenu()
}
