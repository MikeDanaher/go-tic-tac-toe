package tic_tac_toe

import (
	"fmt"
	"strings"
)

type Board struct {
	rowSize int
	cells   []string
}

func NewBoard(s int) *Board {
	cells := make([]string, s*s)
	return &Board{s, cells}
}

func (b *Board) MakeMove(cell int, symbol string) {
	b.cells[cell-1] = symbol
}

func (b *Board) Cells() []string {
	return b.cells
}

func (b *Board) AvailableCells() []int {
	var available []int
	for i, v := range b.cells {
		if v == "" {
			available = append(available, i+1)
		}
	}
	return available
}

func (b *Board) IsFull() bool {
	for _, i := range b.cells {
		if i == "" {
			return false
		}
	}
	return true
}

func (b *Board) String() string {
	tempBoard := insertSpaces(b.cells)
	tempRows := createRows(tempBoard, b.rowSize)
	formattedRows := insertVerticalLines(tempRows)
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

func createRows(cells []string, size int) [][]string {
	rows := make([][]string, size)
	counter := 0
	for i := 0; i < size; i++ {
		rows[i] = make([]string, size)
		for j := 0; j < size; j++ {
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
