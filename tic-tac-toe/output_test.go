package tic_tac_toe_test

import (
	"bytes"
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Output", func() {

	It("Writes a string to the given writer", func() {
		mockWriter := bytes.NewBuffer(make([]byte, 10))
		output := NewOutput(mockWriter)

		output.Write("Winner!")
		Expect(mockWriter.String()).Should(ContainSubstring("Winner!"))
	})

})
