package tic_tac_toe

type ComputerAI struct {
	rules *Rules
}

func NewComputerAI(rules *Rules) *ComputerAI {
	return &ComputerAI{rules}
}

func (c *ComputerAI) GetBestMove(board *Board, depth int, currentPlayer string, opponent string, scores map[int]int) int {
	if winner, _ := c.rules.CheckWinner(board); winner {
		return -1
	}

	if board.IsFull() {
		return 0
	}

	for _, cell := range board.AvailableCells() {
		board.MakeMove(cell, currentPlayer)
		scores[cell] = -1 * c.GetBestMove(board, depth-1, opponent, currentPlayer, make(map[int]int))
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
