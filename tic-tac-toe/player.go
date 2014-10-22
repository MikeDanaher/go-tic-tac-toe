package tic_tac_toe

type Player interface {
	AddSymbol(symbol string)
	Symbol() string
	GetMove(board []string, message string) int
}
