package tic_tac_toe

import (
	"fmt"
)

type Game struct {
	Board *Board
}

func NewGame(size int) *Game {
	board := NewBoard(size)
	return &Game{board}
}

func (game *Game) Play(currentPlayer Player, opponent Player) {
	move := currentPlayer.GetMove(game.Board.Cells, ChooseCell)
	game.Board.MakeMove(move, currentPlayer.Symbol())

	if gameOver(game.Board) {
		endGame()
	} else {
		game.Play(opponent, currentPlayer)
	}
}

func gameOver(board *Board) bool {
	for _, i := range board.Cells {
		if i == "" {
			return false
		}
	}
	return true
}

func endGame() {
	fmt.Println("Game Over")
	fmt.Println(MainMenu)
}
