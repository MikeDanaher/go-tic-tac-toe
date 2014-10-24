package tic_tac_toe

import "os"

type Console3x3 struct{}

const (
	MAIN_MENU = "\nTic-Tac-Toe in Go\n" +
		"=====================================================\n" +
		"1. Human vs. Human game\n" +
		"2. Human vs. Computer game\n" +
		"3. Exit\n" +
		"Play the game by entering the number of your choice: "
	CHOOSE_CELL             = "Please select an available cell [1-9]: "
	INVALID_NUMBER          = "Please enter a valid number: "
	INVALID_MOVE            = "Invalid move, try again: "
	CHOOSE_YOUR_SYMBOL      = "Enter your symbol: "
	CHOOSE_OPPONENTS_SYMBOL = "Enter your opponent's symbol: "
	CHOOSE_FIRST_PLAYER     = "\n1. You\n" +
		"2. Opponent\n" +
		"Who will go first? "
	TIE_GAME = "\nIt's a Tie!\n"
)

func (c *Console3x3) ShowWinner(s string) string {
	return "\nPlayer " + s + " wins!\n"
}

func (c *Console3x3) UI() UI {
	return NewConsoleUI(os.Stdin, os.Stdout)
}

func (c *Console3x3) BoardSize() int {
	return 3
}

func (c *Console3x3) WinningLines() [][]int {
	return [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
		[]int{0, 3, 6},
		[]int{1, 4, 7},
		[]int{2, 5, 8},
		[]int{0, 4, 8},
		[]int{2, 4, 6},
	}
}
