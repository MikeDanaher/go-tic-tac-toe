package tic_tac_toe_test

import (
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Computer Player", func() {
	It("Creates a new computer player with a symbol", func() {
		symbol := "o"
		player := NewComputerPlayer()
		player.SetSymbol(symbol)
		Expect(player.Symbol()).To(Equal(symbol))
	})

	It("Gets the best move from the computer player when the board is empty", func() {
		player := NewComputerPlayer()
		player.SetSymbol("o")
		board := NewBoard(3)

		move := player.GetMove(board, CHOOSE_CELL)
		Expect(board.AvailableCells()).To(ContainElement(move))
	})

})
