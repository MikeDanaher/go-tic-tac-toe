package tic_tac_toe

type Game struct {
	ui    UI
	rules *Rules
}

func NewGame(ui UI, rules *Rules) *Game {
	return &Game{ui, rules}
}

func (game *Game) Play(board *Board, currentPlayer Player, opponent Player) {
	game.ui.UpdateBoard(board)
	move := currentPlayer.GetMove(board, opponent.Symbol())
	board.MakeMove(move, currentPlayer.Symbol())

	winner, symbol := game.rules.CheckWinner(board)

	if winner {
		game.ui.EndWithWinner(board, symbol)
	} else if board.IsFull() {
		game.ui.EndWithTie(board)
	} else {
		game.Play(board, opponent, currentPlayer)
	}
}
