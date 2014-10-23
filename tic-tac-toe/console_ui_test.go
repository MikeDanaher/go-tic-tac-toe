package tic_tac_toe_test

import (
	"bytes"
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Console UI", func() {
	var (
		mockWriter *bytes.Buffer
		input      *Input
		output     *Output
		ui         UI
	)

	BeforeEach(func() {
		mockWriter = NewMockWriter()
		input = NewInput(NewMockReader("5\n"))
		output = NewOutput(mockWriter)
		ui = NewConsoleUI(input, output)
	})

	It("Displays a string to the given output", func() {
		message := ChooseCell
		ui.DisplayMessage(message)
		Expect(mockWriter.String()).Should(ContainSubstring(message))
	})

	It("Gets a numerical response from the user", func() {
		move := ui.GetNumber()
		Expect(move).To(Equal(5))
	})

})
