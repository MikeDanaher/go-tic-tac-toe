package tic_tac_toe

type Player interface {
	SetSymbol(symbol string)
	Symbol() string
	GetMove(*Board, string, string) int
}
