package tic_tac_toe

type HumanPlayer struct {
	ui     UI
	rules  *Rules
	symbol string
}

func NewHumanPlayer(ui UI, rules *Rules) Player {
	return &HumanPlayer{ui: ui, rules: rules}
}

func (player *HumanPlayer) SetSymbol(symbol string) {
	player.symbol = symbol
}

func (player *HumanPlayer) Symbol() string {
	return player.symbol
}

func (player *HumanPlayer) GetMove(board *Board, opponent string) int {
	msg := ChooseCell(board.AvailableCells())
	return getValidMove(player.ui, player.rules, board, msg)
}

func getValidMove(ui UI, rules *Rules, board *Board, message string) int {
	ui.DisplayMessage(message)
	move := ui.GetNumber()

	if rules.ValidMove(move, board) {
		return move
	}
	return getValidMove(ui, rules, board, INVALID_MOVE)
}
