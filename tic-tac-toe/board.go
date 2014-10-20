package tic_tac_toe

type Board struct {
	Size  int
	Cells []string
}

func NewBoard(s int) *Board {
	size := s * s
	cells := make([]string, size)
	return &Board{size, cells}
}

func (b *Board) MakeMove(cell int, symbol string) {
	b.Cells[cell] = symbol
}
