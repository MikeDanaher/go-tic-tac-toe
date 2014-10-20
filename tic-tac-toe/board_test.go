package tic_tac_toe_test

import (
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board", func() {
	var board *Board

	BeforeEach(func() {
		board = NewBoard(3)
	})

	It("Creates a new board with size 3", func() {
		Expect(board.Size).To(Equal(9))
	})

	It("Fills a new board with empty cells", func() {
		cells := make([]string, 9)
		Expect(board.Cells).To(Equal(cells))
	})

	It("Fills a cell on the board with a symbol", func() {
		board.MakeMove(2, "x")
		Expect(board.Cells[2]).To(Equal("x"))
	})
})
