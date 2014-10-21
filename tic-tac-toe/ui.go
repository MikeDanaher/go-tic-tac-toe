package tic_tac_toe

type UI interface {
	PromptForInput(s string)
	GetNumber() int
}
