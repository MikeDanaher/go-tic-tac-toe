package tic_tac_toe_test

import (
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Human Player", func() {
	var ui UI

	BeforeEach(func() {
		mockReader := NewMockReader("5")
		mockWriter := NewMockWriter()
		ui = NewConsoleUI(NewInput(mockReader), NewOutput(mockWriter))
	})

	It("Creates a new human player with a symbol", func() {
		symbol := "x"
		player := NewHumanPlayer(symbol, ui)

		Expect(player.Symbol).To(Equal("x"))
	})

	It("Gets the next move from the human player", func() {
		player := NewHumanPlayer("x", ui)
		board := []string{"", "x", "o", "", "", "", "", "o", "x"}
		availableCells := []int{0, 3, 4, 5, 6}

		move := player.GetMove(board)

		Expect(availableCells).To(ContainElement(move))
	})

})
