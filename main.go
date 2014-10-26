package main

import (
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
)

func main() {
	setup := NewConsoleSetup(new(Config3x3))
	setup.Run()
	main()
}
