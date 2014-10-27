package tic_tac_toe

type ComputerPlayer struct {
	Rules  *Rules
	symbol string
}

func NewComputerPlayer(rules *Rules) Player {
	return &ComputerPlayer{Rules: rules}
}

func (player *ComputerPlayer) GetMove(board *Board, opponent string, message string) int {
	scores := make(map[int]int)
	initialDepth := 0
	return getBestMove(board, initialDepth, player.Symbol(), opponent, scores, player.Rules)
}

func (player *ComputerPlayer) SetSymbol(symbol string) {
	player.symbol = symbol
}

func (player *ComputerPlayer) Symbol() string {
	return player.symbol
}

func getBestMove(board *Board, depth int, currentPlayer string, opponent string, scores map[int]int, rules *Rules) int {
	if winner, _ := rules.Winner(board); winner {
		return -1
	}

	if board.IsFull() {
		return 0
	}

	for _, cell := range board.AvailableCells() {
		board.MakeMove(cell, currentPlayer)
		scores[cell] = -1 * getBestMove(board, depth-1, opponent, currentPlayer, make(map[int]int), rules)
		board.Remove(cell)
	}

	bestMove, highestScore := maxScore(scores)

	if depth == 0 {
		return bestMove
	}
	return highestScore
}

func maxScore(scores map[int]int) (cell int, score int) {
	cell = -1
	score = -2
	for key, value := range scores {
		if value > score {
			cell = key
			score = value
		}
	}
	return
}
