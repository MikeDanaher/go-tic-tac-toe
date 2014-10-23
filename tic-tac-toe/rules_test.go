package tic_tac_toe_test

import (
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rules", func() {
	var config Config
	var board *Board
	var rules *Rules

	BeforeEach(func() {
		config = new(Config3x3)
		board = NewBoard(config.BoardSize())
		rules = NewRules(config)
	})

	It("Knows if a move is valid", func() {
		board.MakeMove(2, "x")
		Expect(rules.ValidMove(3, board)).To(BeTrue())
	})

	It("Knows if a move is invalid", func() {
		board.MakeMove(2, "x")
		Expect(rules.ValidMove(2, board)).To(BeFalse())
	})

	It("Knows if there is a winner", func() {
		board.MakeMove(0, "x")
		board.MakeMove(1, "x")
		board.MakeMove(2, "x")

		winner, symbol := rules.Winner(board)
		Expect(winner).To(BeTrue())
		Expect(symbol).To(Equal("x"))
	})

	It("Knows if there is not a winner", func() {
		board.MakeMove(0, "x")
		board.MakeMove(2, "x")

		winner, symbol := rules.Winner(board)
		Expect(winner).To(BeFalse())
		Expect(symbol).To(Equal(""))
	})
})
