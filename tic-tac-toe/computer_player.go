package tic_tac_toe

type ComputerPlayer struct {
	rules  *Rules
	symbol string
}

func NewComputerPlayer(rules *Rules) Player {
	return &ComputerPlayer{rules: rules}
}

func (player *ComputerPlayer) GetMove(board *Board, opponent string) int {
	ai := NewComputerAI(player.rules)
	scores := make(map[int]int)
	initialDepth := 0
	return ai.GetBestMove(board, initialDepth, player.Symbol(), opponent, scores)
}

func (player *ComputerPlayer) SetSymbol(symbol string) {
	player.symbol = symbol
}

func (player *ComputerPlayer) Symbol() string {
	return player.symbol
}
