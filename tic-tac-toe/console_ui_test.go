package tic_tac_toe_test

import (
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Console UI", func() {

	It("Prompts the user with a message", func() {
		message := "Please select an available cell: "
		input := NewInput(NewMockReader("5"))

		mockWriter := NewMockWriter()
		output := NewOutput(mockWriter)
		ui := NewConsoleUI(input, output)

		ui.PromptForInput(message)

		Expect(mockWriter.String()).Should(ContainSubstring(message))
	})

	It("Gets the response from the user", func() {
		input := NewInput(NewMockReader("5"))
		output := NewOutput(NewMockWriter())
		ui := NewConsoleUI(input, output)

		move := ui.GetNumber()
		Expect(move).To(Equal(5))
	})

})
