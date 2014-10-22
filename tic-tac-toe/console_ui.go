package tic_tac_toe

import (
	"strconv"
)

type ConsoleUI struct {
	in  *Input
	out *Output
}

const (
	MainMenu string = "Tic-Tac-Toe in Go\n" +
		"=====================================================\n" +
		"Play the game by entering the number of your choice:" +
		"1. Human vs. Human game" +
		"2. Human vs. Computer game" +
		"3. Exit"
	ChooseCell    string = "Please select an available cell [1-9]: "
	InvalidNumber string = "Please enter a valid number: "
)

func NewConsoleUI(in *Input, out *Output) *ConsoleUI {
	return &ConsoleUI{in, out}
}

func (c *ConsoleUI) PromptForInput(message string) {
	c.out.Write(message)
}

func (c *ConsoleUI) GetNumber() int {
	input, err := strconv.Atoi(c.in.Read())

	if err != nil {
		c.PromptForInput(InvalidNumber)
		c.GetNumber()
	}

	return input
}
