package tic_tac_toe_test

import (
	"bytes"
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Human Player", func() {
	var (
		ui         UI
		mockReader bytes.Buffer
		mockWriter bytes.Buffer
		board      *Board
		rules      *Rules
		player     Player
		symbol     string
		opponent   string
	)

	BeforeEach(func() {
		mockReader.Reset()
		ui = NewConsoleUI(&mockReader, &mockWriter)
		board = NewBoard(3)
		rules = new(Rules)
		symbol = "x"
		opponent = "o"
		player = MakeHumanPlayer(symbol, ui, rules)
	})

	It("Creates a new human player with a symbol", func() {
		Expect(player.Symbol()).To(Equal(symbol))
	})

	It("Gets the next move from the human player", func() {
		BuildBoard(board, []int{2, 7}, symbol)
		BuildBoard(board, []int{3}, opponent)
		mockReader.WriteString("5\n")

		move := player.GetMove(board, opponent)

		Expect(board.AvailableCells()).To(ContainElement(move))
	})

	It("Prompts again if the given move was invalid", func() {
		BuildBoard(board, []int{2, 7}, symbol)
		BuildBoard(board, []int{4, 3}, opponent)
		mockReader.WriteString("4\n1\n")

		move := player.GetMove(board, opponent)

		Expect(move).To(Equal(1))
		Expect(mockWriter.String()).Should(ContainSubstring(INVALID_MOVE))
	})
})
