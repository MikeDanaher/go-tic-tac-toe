package tic_tac_toe

type MockPlayer struct {
	symbol string
	moves  []int
}

func NewMockPlayer(moves []int) Player {
	return &MockPlayer{moves: moves}
}

func (mp *MockPlayer) SetSymbol(s string) {
	mp.symbol = s
}

func (mp *MockPlayer) Symbol() string {
	return mp.symbol
}

func (mp *MockPlayer) GetMove(board *Board, s string) int {
	var x int
	x, mp.moves = mp.moves[len(mp.moves)-1], mp.moves[:len(mp.moves)-1]
	return x
}
