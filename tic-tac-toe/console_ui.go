package tic_tac_toe

import (
	"strconv"
)

type ConsoleUI struct {
	in  *Input
	out *Output
}

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

func NewConsoleUI(in *Input, out *Output) UI {
	return &ConsoleUI{in, out}
}

func (c *ConsoleUI) DisplayMessage(message string) {
	c.out.Write(message)
}

func (c *ConsoleUI) GetNumber() int {
	input, err := strconv.Atoi(c.in.Read())

	if err != nil {
		c.DisplayMessage(InvalidNumber)
		return c.GetNumber()
	}

	return input
}
