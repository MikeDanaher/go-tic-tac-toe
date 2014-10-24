package tic_tac_toe_test

import (
	"bytes"
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func makePlayer(symbol string, ui UI, rules *Rules) Player {
	player := NewHumanPlayer(ui, rules)
	player.SetSymbol(symbol)
	return player
}

var _ = Describe("Human Player", func() {
	var (
		ui         UI
		mockReader bytes.Buffer
		mockWriter bytes.Buffer
		board      *Board
		rules      *Rules
	)

	BeforeEach(func() {
		mockReader.Reset()
		ui = NewConsoleUI(&mockReader, &mockWriter)
		rules = new(Rules)
		board = NewBoard(3)
	})

	It("Creates a new human player with a symbol", func() {
		symbol := "x"
		player := makePlayer(symbol, ui, rules)
		Expect(player.Symbol()).To(Equal(symbol))
	})

	It("Gets the next move from the human player", func() {
		player := makePlayer("x", ui, rules)
		board.MakeMove(2, "x")
		board.MakeMove(3, "o")
		board.MakeMove(7, "x")
		availableCells := []int{0, 1, 4, 5, 6, 8}

		moves := "5\n1\n"
		mockReader.WriteString(moves)

		move := player.GetMove(board, ChooseCell)
		Expect(availableCells).To(ContainElement(move))
	})

	It("Prompts again if the given move was invalid", func() {
		player := makePlayer("x", ui, rules)
		board.MakeMove(2, "x")
		board.MakeMove(4, "o")
		board.MakeMove(3, "o")
		board.MakeMove(7, "x")

		moves := "5\n1\n"
		mockReader.WriteString(moves)

		move := player.GetMove(board, ChooseCell)
		Expect(move).To(Equal(0))
		Expect(mockWriter.String()).Should(ContainSubstring("Invalid move, try again: "))
	})
})
