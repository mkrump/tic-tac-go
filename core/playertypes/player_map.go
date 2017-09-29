package playertypes

import (
	"fmt"
)

type PlayerMap struct {
	Players map[int]Player
}

type PlayerSymbolNotFoundError struct {
	val int
}

func (e *PlayerSymbolNotFoundError) Error() string {
	return fmt.Sprintf("%s does not have a player associated", string(e.val))
}

func (players PlayerMap) PlayerSymbol(player int) (string, error) {
	if player, ok := players.Players[player]; ok {
		return player.Symbol(), nil
	}
	return "", &PlayerSymbolNotFoundError{}
}

func (players PlayerMap) Player(player int) Player {
	return players.Players[player]
}
