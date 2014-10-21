package tic_tac_toe

type HumanPlayer struct {
	Symbol string
	ui     UI
}

func NewHumanPlayer(s string, ui UI) *HumanPlayer {
	return &HumanPlayer{s, ui}
}

func (player *HumanPlayer) GetMove(board []string) int {
	//player.ui.PromptForInput()
	//move = player.ui.GetInput()
	move := 3

	if validMove(board, move) {
		return move
	}
	return player.GetMove(board)
}

// Maybe move to a utilities class?
func validMove(board []string, move int) bool {
	if board[move] == "" {
		return true
	}
	return false
}
