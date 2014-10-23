package main

import (
	"bytes"
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	"os"
)

func main() {
	console.DisplayMessage(MainMenu)

	rules := NewRules(new(Console3x3))
	board := NewBoard(rules.BoardSize)

	input := NewInput(bytes.NewBuffer(os.Stdin))
	output := NewOutput(bytes.NewBuffer(os.Stdout))
	console := NewConsoleUI(input, output)

	player1 := NewHumanPlayer(console, rules)
	player1.SetSymbol("x")
	player2 := NewHumanPlayer(console, rules)
	player2.SetSymbol("o")

	game := NewGame(console, rules)
	game.Play(board, player1, player2)
}
