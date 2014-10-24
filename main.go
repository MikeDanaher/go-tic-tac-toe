package main

import (
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
)

func main() {
	setup := NewSetup(new(Console3x3))
	setup.Run()
}
