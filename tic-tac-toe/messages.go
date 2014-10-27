package tic_tac_toe

const (
	MAIN_MENU = "\nTIC-TAC-TOE in GO\n" +
		"=====================================================\n" +
		"1. Human vs. Human game\n" +
		"2. Human vs. Computer game\n" +
		"3. Computer vs. Computer game\n" +
		"4. Exit\n" +
		"Play the game by entering the number of your choice: "
	CHOOSE_CELL             = "Please select an available cell [1-9]: "
	INVALID_NUMBER          = "Please enter a valid number: "
	INVALID_MOVE            = "Invalid move, try again: "
	CHOOSE_YOUR_SYMBOL      = "Enter your symbol: "
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
