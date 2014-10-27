package tic_tac_toe

import "fmt"

const (
	MAIN_MENU = "\nTIC-TAC-TOE in GO\n" +
		"=====================================================\n" +
		"1. Human vs. Human game\n" +
		"2. Human vs. Computer game\n" +
		"3. Computer vs. Computer game\n" +
		"4. Exit\n" +
		"Choose a game by entering the number: "
	DESCRIBE_GAME           = "\nOn your turn, place your symbol in a cell by entering the number that corresponds to the cell you want.\nThe cells are numbered as follows: "
	CELL_NUMBERS            = "\n 1 | 2 | 3 \n-----------\n 4 | 5 | 6 \n-----------\n 7 | 8 | 9 \n\n"
	INVALID_NUMBER          = "Please enter a valid number: "
	INVALID_MOVE            = "Invalid move, try again: "
	CHOOSE_YOUR_SYMBOL      = "\nEnter your symbol: "
	CHOOSE_OPPONENTS_SYMBOL = "Enter your opponent's symbol: "
	TIE_GAME                = "\n*** It's a Tie! ***\n"
	PLAY_AGAIN              = "\nPlay Again?\n"
	HUMAN_v_HUMAN           = "HvH"
	HUMAN_v_COMPUTER        = "HvC"
	COMPUTER_v_COMPUTER     = "CvC"
	CHOOSE_FIRST_PLAYER     = "\n1. You\n" +
		"2. Opponent\n" +
		"Who will go first? "
)

func PrintWinner(p string) string {
	return "\n*** Player " + p + " wins! ***\n"
}

func ChooseCell(availableCells []int) string {
	return fmt.Sprintf("Please select an available cell %v: ", availableCells)
}
