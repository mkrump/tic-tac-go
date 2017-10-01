package menus

import "github.com/sc2nomore/tic-tac-go/core"

type StartupMenu interface {
	PlayerTypePrompt() (string, error)
	PlayerSymbolPrompt() (string, error)
	SelectPlayerType(string, string) (core.Player, error)
}
