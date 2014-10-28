package tic_tac_toe_test

import (
	"bytes"
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Console UI", func() {
	var (
		ui         UI
		mockWriter bytes.Buffer
		mockReader bytes.Buffer
	)

	BeforeEach(func() {
		ui = NewConsoleUI(&mockReader, &mockWriter)
	})

	It("Displays a string to the given output", func() {
		message := DESCRIBE_GAME
		ui.DisplayMessage(message)
		Expect(mockWriter.String()).Should(ContainSubstring(message))
	})

	It("Gets a numerical response from the user", func() {
		mockReader.WriteString("5\n")
		move := ui.GetNumber()
		Expect(move).To(Equal(5))
	})

	It("Gets a string response from the user", func() {
		mockReader.WriteString("symbol\n")
		move := ui.GetString()
		Expect(move).To(Equal("symbol"))
	})

	It("Prints a board", func() {
		board := NewBoard(3)
		ui.UpdateBoard(board)
		Expect(mockWriter.String()).To(ContainSubstring(board.String()))
	})

	It("Ends the game with a winner", func() {
		board := NewBoard(3)
		symbol := "x"
		ui.EndWithWinner(board, symbol)
		Expect(mockWriter.String()).To(ContainSubstring(PrintWinner(symbol)))
	})

	It("Ends the game with a tie", func() {
		board := NewBoard(3)
		ui.EndWithTie(board)
		Expect(mockWriter.String()).To(ContainSubstring(TIE_GAME))
	})
})
