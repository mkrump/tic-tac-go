package games

import "github.com/sc2nomore/tic-tac-go/core"

type Game struct {
	Players core.PlayerMap
	Board   core.Board
	Rules   core.Rules
}

func (game Game) GameBoard() core.Board {
	return game.Board
}

func (game Game) GamePlayers() core.PlayerMap {
	return game.Players
}

func (game Game) GetMove() interface{} {
	board := game.GameBoard()
	activePlayer := game.ActivePlayer()
	activePlayerNumber := game.Rules.ActivePlayerNumber(board)
	return activePlayer.Move(game.Board, activePlayerNumber)
}

//MakeGame constructor for game struct
func MakeGame(board core.Board, players core.PlayerMap, rules core.Rules) Game {
	return Game{
		Board:   board,
		Players: players,
		Rules:   rules,
	}
}

// IsWin checks if the game is a win for the current player
// and returns a boolean
func (game Game) IsWin() bool {
	board := game.GameBoard()
	InActivePlayerNumber := game.Rules.InActivePlayerNumber(board)
	return game.Rules.IsWin(game.Board,	InActivePlayerNumber)
}

// IsWin checks if the game is a tie and returns a boolean
func (game Game) IsTie() bool {
	return game.Rules.IsTie(game.Board)
}

// MakeMove attempts make a move and updates the TTTBoard state
// if valid and returns an error if the move is invalid
func (game Game) MakeMove(move int) error {
	board := game.GameBoard()
	activePlayerNumber := game.Rules.ActivePlayerNumber(board)
	return game.Board.MakeMove(move, activePlayerNumber)
}

func (game Game) InActivePlayer() core.Player {
	board := game.GameBoard()
	InActivePlayerNumber := game.Rules.InActivePlayerNumber(board)
	return game.Players.Player(InActivePlayerNumber)
}

func (game Game) ActivePlayer() core.Player {
	board := game.GameBoard()
	activePlayerNumber := game.Rules.ActivePlayerNumber(board)
	return game.Players.Player(activePlayerNumber)
}

func (game Game) InActivePlayerMarker() string {
	player := game.InActivePlayer()
	return player.Symbol()
}

func (game Game) ActivePlayerMarker() string {
	player := game.ActivePlayer()
	return player.Symbol()
}

func MakePlayers(player1 core.Player, player2 core.Player) core.PlayerMap {
	return core.PlayerMap{
		Players: map[int]core.Player{
			-1: player1,
			1:  player2,
		},
	}
}
