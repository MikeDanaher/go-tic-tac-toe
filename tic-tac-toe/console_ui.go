package tic_tac_toe

import (
	"strconv"
)

type ConsoleUI struct {
	in  *Input
	out *Output
}

func NewConsoleUI(in *Input, out *Output) *ConsoleUI {
	return &ConsoleUI{in, out}
}

func (c *ConsoleUI) PromptForInput(message string) {
	c.out.Write(message)
}

func (c *ConsoleUI) GetNumber() int {
	input, err := strconv.Atoi(c.in.Read())

	if err != nil {
		c.PromptForInput("Please enter a valid number: ")
		c.GetNumber()
	}

	return input
}
