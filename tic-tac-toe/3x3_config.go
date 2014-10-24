package tic_tac_toe

type Config3x3 struct{}

const (
	MainMenu = "Tic-Tac-Toe in Go\n" +
		"=====================================================\n" +
		"Play the game by entering the number of your choice:\n" +
		"1. Human vs. Human game\n" +
		"2. Human vs. Computer game\n" +
		"3. Exit\n"
	ChooseCell    = "Please select an available cell [1-9]: "
	InvalidNumber = "Please enter a valid number: "
	InvalidMove   = "Invalid move, try again: "
)

func (c *Config3x3) BoardSize() int {
	return 3
}

func (c *Config3x3) WinningLines() [][]int {
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
