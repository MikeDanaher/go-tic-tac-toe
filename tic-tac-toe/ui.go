package tic_tac_toe

type UI interface {
	DisplayMessage(string)
	GetNumber() int
	GetString() string
	PrintBoard(*Board)
}
