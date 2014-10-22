package tic_tac_toe

type HumanPlayer struct {
	ui     UI
	symbol string
}

func NewHumanPlayer(ui UI) Player {
	player := new(HumanPlayer)
	player.ui = ui
	return player
}

const (
	InvalidMove = "Invalid move, try again: "
)

func (player *HumanPlayer) GetMove(board []string, message string) int {
	player.ui.DisplayMessage(player.ui.FormatBoard(board))
	player.ui.DisplayMessage(message)
	move := player.ui.GetNumber() - 1

	if validMove(board, move) {
		return move
	}
	return player.GetMove(board, InvalidMove)
}

func (player *HumanPlayer) AddSymbol(symbol string) {
	player.symbol = symbol
}

func (player *HumanPlayer) Symbol() string {
	return player.symbol
}

// Maybe move to a utilities class?
func validMove(board []string, move int) bool {
	if board[move] == "" {
		return true
	}
	return false
}
