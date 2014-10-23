package tic_tac_toe

type Player interface {
	SetSymbol(symbol string)
	Symbol() string
	GetMove(board *Board, message string) int
}
