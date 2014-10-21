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

func (b *Board) AvailableCells() []int {
	var available []int
	for i, symbol := range b.Cells {
		if symbol == "" {
			available = append(available, i)
		}
	}
	return available
}
