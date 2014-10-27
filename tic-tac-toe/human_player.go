package tic_tac_toe

type HumanPlayer struct {
	ui     UI
	rules  *Rules
	symbol string
}

func NewHumanPlayer(ui UI, rules *Rules) Player {
	return &HumanPlayer{ui: ui, rules: rules}
}

func (player *HumanPlayer) GetMove(board *Board, opponent string, message string) int {
	player.ui.DisplayMessage(message)
	move := player.ui.GetNumber()

	if player.rules.ValidMove(move, board) {
		return move
	}
	return player.GetMove(board, opponent, INVALID_MOVE)
}

func (player *HumanPlayer) SetSymbol(symbol string) {
	player.symbol = symbol
}

func (player *HumanPlayer) Symbol() string {
	return player.symbol
}
