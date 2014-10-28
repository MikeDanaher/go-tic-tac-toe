package tic_tac_toe

type UI interface {
	DisplayMessage(string)
	GetNumber() int
	GetString() string
	UpdateBoard(*Board)
	EndWithWinner(*Board, string)
	EndWithTie(*Board)
}
