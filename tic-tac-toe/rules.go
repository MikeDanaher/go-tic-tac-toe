package tic_tac_toe

type Rules struct {
	winningLines [][]int
}

func NewRules(winningLines [][]int) *Rules {
	return &Rules{winningLines}
}

func (r *Rules) ValidMove(move int, board *Board) bool {
	for _, v := range board.AvailableCells() {
		if v == move {
			return true
		}
	}
	return false
}

func (r *Rules) Winner(board *Board) (winner bool, symbol string) {
	var line []string
	for _, i := range r.winningLines {
		line = mapSymbols(i, board.Cells())
		if uniqueSymbols(line) {
			winner = true
			symbol = line[0]
			return
		}
	}
	winner = false
	symbol = ""
	return
}

func mapSymbols(winningLine []int, cells []string) []string {
	var container []string
	for _, n := range winningLine {
		container = append(container, cells[n])
	}
	return container
}

func uniqueSymbols(line []string) bool {
	s := line[0]
	for n := range line {
		if s != line[n] || s == "" {
			return false
		}
	}
	return true
}
