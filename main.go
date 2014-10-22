package main

import (
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	"os"
)

func main() {
	input := NewInput(os.Stdin)
	output := NewOutput(os.Stdout)
	console := NewConsoleUI(input, output)

	player1 := NewHumanPlayer(console)
	player2 := NewHumanPlayer(console)

	console.DisplayMessage(MainMenu)

	player1.AddSymbol("x")
	player2.AddSymbol("o")

	game := NewGame(3)

	game.Play(player1, player2)
}
