package tic_tac_toe

type HumanPlayer struct {
	Symbol string
	ui     UI
}

func NewHumanPlayer(s string, ui UI) *HumanPlayer {
	return &HumanPlayer{s, ui}
}

const (
	InvalidMove = "Invalid move, try again: "
)

func (player *HumanPlayer) GetMove(board []string, message string) int {
	player.ui.PromptForInput(message)
	move := player.ui.GetNumber()

	if validMove(board, move) {
		return move
	}
	return player.GetMove(board, InvalidMove)
}

// Maybe move to a utilities class?
func validMove(board []string, move int) bool {
	if board[move] == "" {
		return true
	}
	return false
}
