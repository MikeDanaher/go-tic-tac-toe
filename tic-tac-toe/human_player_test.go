package tic_tac_toe_test

import (
	"bytes"
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func makePlayer(symbol string, ui UI) Player {
	player := NewHumanPlayer(ui)
	player.AddSymbol(symbol)
	return player
}

var _ = Describe("Human Player", func() {
	var (
		ui         UI
		mockReader *bytes.Buffer
		mockWriter *bytes.Buffer
	)

	BeforeEach(func() {
		mockReader = NewMockReader("5\n")
		mockWriter = NewMockWriter()
		ui = NewConsoleUI(NewInput(mockReader), NewOutput(mockWriter))
	})

	It("Creates a new human player with a symbol", func() {
		symbol := "x"
		player := makePlayer(symbol, ui)
		Expect(player.Symbol()).To(Equal(symbol))
	})

	It("Gets the next move from the human player", func() {
		player := makePlayer("x", ui)
		board := []string{"", "x", "o", "", "", "", "", "o", "x"}
		availableCells := []int{0, 3, 4, 5, 6}

		move := player.GetMove(board, ChooseCell)
		Expect(availableCells).To(ContainElement(move))
	})

})
