package tic_tac_toe_test

import (
	"bytes"
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game", func() {
	var (
		ui         UI
		config     Config
		mockWriter bytes.Buffer
		mockReader bytes.Buffer
		board      *Board
		rules      *Rules
		p1         Player
		p1Symbol   string
		p2         Player
		p2Symbol   string
		game       *Game
	)

	BeforeEach(func() {
		ui = NewConsoleUI(&mockReader, &mockWriter)
		config = new(Config3x3)
		board = NewBoard(config.BoardSize())
		rules = NewRules(config.WinningLines())
		game = NewGame(ui, rules)
		p1Symbol = "x"
		p2Symbol = "o"
	})

	It("Ends in a tie if the board fills up", func() {
		p1Moves := []int{7, 6, 2, 9, 5}
		p1 = NewMockPlayer(p1Moves)
		p2Moves := []int{4, 8, 3, 1}
		p2 = NewMockPlayer(p2Moves)
		p1.SetSymbol(p1Symbol)
		p2.SetSymbol(p2Symbol)

		game.Play(board, p1, p2)

		Expect(mockWriter.String()).To(ContainSubstring(TIE_GAME))
	})

	It("Ends with a winner if a player gets three in a row", func() {
		p1Moves := []int{6, 3, 9, 5}
		p1 = NewMockPlayer(p1Moves)
		p2Moves := []int{7, 2, 1}
		p2 = NewMockPlayer(p2Moves)
		p1.SetSymbol(p1Symbol)
		p2.SetSymbol(p2Symbol)

		game.Play(board, p1, p2)

		Expect(mockWriter.String()).To(ContainSubstring(PrintWinner(p1Symbol)))
	})
})
