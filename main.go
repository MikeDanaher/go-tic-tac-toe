package main

import (
	"bufio"
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	"os"
)

func main() {
	config := new(Config3x3)
	rules := NewRules(config)
	board := NewBoard(config.BoardSize())

	input := NewInput(bufio.NewReader(os.Stdin))
	output := NewOutput(os.Stdout)
	console := NewConsoleUI(input, output)

	console.DisplayMessage(MainMenu)

	player1 := NewHumanPlayer(console, rules)
	player1.SetSymbol("x")
	player2 := NewHumanPlayer(console, rules)
	player2.SetSymbol("o")

	game := NewGame(console, rules)
	game.Play(board, player1, player2)
}
