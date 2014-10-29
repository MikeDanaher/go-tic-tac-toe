package tic_tac_toe

type Config interface {
	BoardSize() int
	WinningLines() [][]int
	Input() Reader
	Output() Writer
}
