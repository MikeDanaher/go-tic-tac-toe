package tic_tac_toe_test

import (
	. "../tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board", func() {
	var board *Board

	BeforeEach(func() {
		board = NewBoard(3)
	})

	It("Creates a blank 3x3 board", func() {
		Expect(len(board.Cells())).To(Equal(9))
	})

	It("Fills a cell on the board with a symbol", func() {
		board.MakeMove(2, "x")
		Expect(board.Cells()[1]).To(Equal("x"))
	})

	It("Removes a symbol from the board", func() {
		board.MakeMove(2, "x")
		board.Remove(2)
		Expect(board.Cells()[1]).To(Equal(""))
	})

	It("Gets the available cells on the board", func() {
		board.MakeMove(3, "o")
		board.MakeMove(6, "x")
		board.MakeMove(8, "0")
		Expect(board.AvailableCells()).To(Equal([]int{1, 2, 4, 5, 7, 9}))
	})

	It("Knows if the board is not full", func() {
		board.MakeMove(2, "x")
		Expect(board.IsFull()).To(BeFalse())
	})

	It("Knows if the board is full", func() {
		board.MakeMove(1, "o")
		board.MakeMove(2, "x")
		board.MakeMove(3, "o")
		board.MakeMove(4, "x")
		board.MakeMove(5, "o")
		board.MakeMove(6, "x")
		board.MakeMove(7, "o")
		board.MakeMove(8, "x")
		board.MakeMove(9, "x")
		Expect(board.IsFull()).To(BeTrue())
	})

	It("Gets the board as a string", func() {
		board.MakeMove(1, "o")
		board.MakeMove(4, "x")
		board.MakeMove(5, "o")
		board.MakeMove(8, "x")
		Expect(board.String()).To(Equal("\n o |   |   \n-----------\n x | o |   \n-----------\n   | x |   \n\n"))
	})
})
