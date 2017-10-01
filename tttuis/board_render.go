package tttuis

import "github.com/sc2nomore/tic-tac-go/core"

type BoardRender interface {
	RenderBoard(board core.Board, players core.PlayerMapper) string
}
