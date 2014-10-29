package main

import (
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	"os"
)

func main() {
	config := NewConfig3x3(os.Stdin, os.Stdout)
	setup := NewConsoleSetup(config)
	setup.Run()
	main()
}
