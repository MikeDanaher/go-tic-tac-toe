package tic_tac_toe_test

import (
	"bytes"
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Input", func() {

	It("Reads the input from the given reader", func() {
		mockReader := bytes.NewBufferString("Hello\n")
		input := NewInput(mockReader)
		text, _ := input.Read()
		Expect(text).To(Equal("Hello\n"))
	})

	It("Has an error if the input is incorrect", func() {
		mockReader := bytes.NewBufferString("")
		input := NewInput(mockReader)
		_, err := input.Read()
		Expect(err).Should(HaveOccurred())
	})
})
