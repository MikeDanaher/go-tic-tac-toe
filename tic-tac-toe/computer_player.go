package tic_tac_toe

import "math/rand"

type ComputerPlayer struct {
	symbol string
}

func NewComputerPlayer() Player {
	return &ComputerPlayer{}
}

func (player *ComputerPlayer) GetMove(board *Board, message string) int {
	randomMove := rand.Intn(len(board.AvailableCells()))
	return board.AvailableCells()[randomMove]
}

func (player *ComputerPlayer) SetSymbol(symbol string) {
	player.symbol = symbol
}

func (player *ComputerPlayer) Symbol() string {
	return player.symbol
}
