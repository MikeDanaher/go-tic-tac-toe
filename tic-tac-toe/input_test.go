package tic_tac_toe_test

import (
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GameInput", func() {

	It("Reads the input from the given reader", func() {
		mockReader := NewMockReader("Hello\n")
		input := NewInput(mockReader)
		text := input.Read()
		Expect(text).To(Equal("Hello\n"))
	})

})
