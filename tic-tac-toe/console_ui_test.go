package tic_tac_toe_test

import (
	"bytes"
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Console UI", func() {
	var ui UI
	var mockWriter bytes.Buffer
	var mockReader bytes.Buffer

	BeforeEach(func() {
		ui = NewConsoleUI(&mockReader, &mockWriter)
	})

	It("Displays a string to the given output", func() {
		message := CHOOSE_CELL
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
		ui.PrintBoard(board)
		Expect(mockWriter.String()).Should(ContainSubstring(board.String()))
	})

})
