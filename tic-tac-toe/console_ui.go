package tic_tac_toe

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type ConsoleUI struct {
	in  *Input
	out *Output
}

const (
	MainMenu string = "Tic-Tac-Toe in Go\n" +
		"=====================================================\n" +
		"Play the game by entering the number of your choice:\n" +
		"1. Human vs. Human game\n" +
		"2. Human vs. Computer game\n" +
		"3. Exit\n"
	ChooseCell    string = "Please select an available cell [1-9]: "
	InvalidNumber string = "Please enter a valid number: "
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

func (c *ConsoleUI) FormatBoard(cells []string) string {
	board := insertSpaces(cells)
	boardRows := createRows(board)
	formattedRows := insertVerticalLines(boardRows)
	formattedBoard := insertHorizontalLines(formattedRows)
	return fmt.Sprintf("\n%v\n\n", formattedBoard)
}

func insertSpaces(cells []string) []string {
	var container []string
	var symbol string
	for _, s := range cells {
		if s == "" {
			symbol = "   "
		} else {
			symbol = fmt.Sprintf(" %v ", s)
		}
		container = append(container, symbol)
	}
	return container
}

func createRows(cells []string) [][]string {
	rowLength := int(math.Sqrt(float64(len(cells))))
	rows := make([][]string, rowLength)
	counter := 0
	for i := 0; i < rowLength; i++ {
		rows[i] = make([]string, rowLength)
		for j := 0; j < rowLength; j++ {
			rows[i][j] = cells[counter]
			counter++
		}
	}
	return rows
}

func insertVerticalLines(board [][]string) []string {
	var rows []string
	for i := range board {
		rows = append(rows, strings.Join(board[i], "|"))
	}
	return rows
}

func insertHorizontalLines(board []string) string {
	return strings.Join(board, "\n-----------\n")
}
