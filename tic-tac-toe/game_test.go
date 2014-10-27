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
		player1    Player
		p1Symbol   string
		player2    Player
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
		player1 = NewComputerPlayer(rules)
		player2 = NewComputerPlayer(rules)
		player1.SetSymbol(p1Symbol)
		player2.SetSymbol(p2Symbol)

		game.Play(board, player1, player2)

		Expect(mockWriter.String()).To(ContainSubstring(TIE_GAME))
	})

	It("Ends with a winner if a player gets three in a row", func() {
		player1 = NewHumanPlayer(ui, rules)
		player2 = NewHumanPlayer(ui, rules)
		player1.SetSymbol(p1Symbol)
		player2.SetSymbol(p2Symbol)
		mockReader.WriteString("5\n2\n9\n3\n1\n")

		game.Play(board, player1, player2)

		Expect(mockWriter.String()).To(ContainSubstring(PrintWinner(p1Symbol)))
	})
})
