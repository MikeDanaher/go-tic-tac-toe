package tic_tac_toe

func MakeHumanPlayer(symbol string, ui UI, rules *Rules) Player {
	player := NewHumanPlayer(ui, rules)
	player.SetSymbol(symbol)
	return player
}

func MakeComputerPlayer(symbol string, rules *Rules) Player {
	player := NewComputerPlayer(rules)
	player.SetSymbol(symbol)
	return player
}

func BuildBoard(board *Board, moves []int, symbol string) {
	for _, move := range moves {
		board.MakeMove(move, symbol)
	}
}
