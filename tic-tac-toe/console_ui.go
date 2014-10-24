package tic_tac_toe

import (
	"io"
	"strconv"
)

type ConsoleUI struct {
	in  Reader
	out Writer
}

func NewConsoleUI(in Reader, out Writer) UI {
	return &ConsoleUI{in, out}
}

func (c *ConsoleUI) DisplayMessage(message string) {
	c.out.WriteString(message)
}

func (c *ConsoleUI) PrintBoard(board *Board) {
	c.out.WriteString(board.String())
}

func (c *ConsoleUI) GetString() string {
	return readLine(c.in)
}

func (c *ConsoleUI) GetNumber() int {
	input, err := strconv.Atoi(readLine(c.in))

	if err != nil {
		c.DisplayMessage(INVALID_NUMBER)
		return c.GetNumber()
	}

	return input
}

func readLine(reader Reader) string {
	var buffer = make([]byte, 1)
	var output string
	for {
		_, err := reader.Read(buffer)
		if buffer[0] == '\n' || err == io.EOF {
			break
		}
		output += string(buffer[0])
	}
	return output
}
