package tic_tac_toe_test

import (
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rules", func() {
	var (
		config Config
		board  *Board
		rules  *Rules
		symbol string
	)

	BeforeEach(func() {
		config = new(Config3x3)
		board = NewBoard(config.BoardSize())
		rules = NewRules(config.WinningLines())
		symbol = "x"
	})

	It("Knows if a move is valid", func() {
		board.MakeMove(2, symbol)
		Expect(rules.ValidMove(3, board)).To(BeTrue())
	})

	It("Knows if a cell has alreayd been taken", func() {
		board.MakeMove(2, symbol)
		Expect(rules.ValidMove(2, board)).To(BeFalse())
	})

	It("Knows if a move is invalid", func() {
		Expect(rules.ValidMove(22, board)).To(BeFalse())
	})

	It("Knows if there is a winner", func() {
		BuildBoard(board, []int{1, 2, 3}, symbol)

		winner, winningSymbol := rules.Winner(board)
		Expect(winner).To(BeTrue())
		Expect(winningSymbol).To(Equal(symbol))
	})

	It("Knows if there is not a winner", func() {
		BuildBoard(board, []int{1, 2, 9}, symbol)

		winner, winningSymbol := rules.Winner(board)
		Expect(winner).To(BeFalse())
		Expect(winningSymbol).To(Equal(""))
	})
})
