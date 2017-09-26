package core

import "github.com/sc2nomore/tic-tac-go/core/playertypes"

//type PlayerSymbolRetreiver interface {
//	PlayerSymbol() string
//}

type PlayerMap struct {
	Players map[int]playertypes.Player
}

func (players PlayerMap) PlayerSymbol(player int) string {
	return players.Players[player].Symbol()
}

func (players PlayerMap) Player(player int) playertypes.Player {
	return players.Players[player]
}
