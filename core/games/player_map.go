package games

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/core"
)

type PlayerMap struct {
	Players map[int]core.Player
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

func (players PlayerMap) Player(player int) core.Player {
	return players.Players[player]
}

func MakePlayers(player1 core.Player, player2 core.Player) core.PlayerMapper {
	return PlayerMap{
		Players: map[int]core.Player{
			-1: player1,
			1:  player2,
		},
	}
}
