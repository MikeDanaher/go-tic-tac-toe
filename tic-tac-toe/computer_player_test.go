package tic_tac_toe_test

import (
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func makeComputerPlayer(symbol string, rules *Rules) Player {
	player := NewComputerPlayer(rules)
	player.SetSymbol(symbol)
	return player
}

func buildBoard(board *Board, moves []int, symbol string) {
	for _, move := range moves {
		board.MakeMove(move, symbol)
	}
}

var _ = Describe("Computer Player", func() {
	var (
		config          Config
		board           *Board
		rules           *Rules
		computer_symbol string
		opponent_symbol string
		player          Player
	)

	BeforeEach(func() {
		config = new(Config3x3)
		board = NewBoard(config.BoardSize())
		rules = NewRules(config.WinningLines())
		computer_symbol = "o"
		opponent_symbol = "x"
		player = makeComputerPlayer(computer_symbol, rules)
	})

	It("Creates a new computer player with a symbol", func() {
		Expect(player.Symbol()).To(Equal(computer_symbol))
	})

	It("Chooses a corner if the opponent has a corner and the center", func() {
		buildBoard(board, []int{5, 9}, opponent_symbol)
		buildBoard(board, []int{1}, computer_symbol)
		corners := []int{3, 7}
		move := player.GetMove(board, opponent_symbol, CHOOSE_CELL)
		Expect(corners).To(ContainElement(move))
	})

	It("Chooses a side if the opponent has opposite corners", func() {
		buildBoard(board, []int{9, 1}, opponent_symbol)
		buildBoard(board, []int{5}, computer_symbol)
		sides := []int{2, 4, 6, 8}
		move := player.GetMove(board, opponent_symbol, CHOOSE_CELL)
		Expect(sides).To(ContainElement(move))
	})

	It("Blocks when the opponent can win", func() {
		buildBoard(board, []int{1, 3, 4, 8}, opponent_symbol)
		buildBoard(board, []int{2, 5, 6}, computer_symbol)
		move := player.GetMove(board, opponent_symbol, CHOOSE_CELL)
		Expect(move).To(Equal(7))
	})

	It("Wins the game when possible", func() {
		buildBoard(board, []int{2, 3, 4, 5}, opponent_symbol)
		buildBoard(board, []int{1, 7, 8}, computer_symbol)
		move := player.GetMove(board, opponent_symbol, CHOOSE_CELL)
		Expect(move).To(Equal(9))
	})
})
