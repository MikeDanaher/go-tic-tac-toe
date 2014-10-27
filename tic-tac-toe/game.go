package tic_tac_toe

type Game struct {
	ui    UI
	rules *Rules
}

func NewGame(ui UI, rules *Rules) *Game {
	return &Game{ui, rules}
}

func (game *Game) Play(board *Board, currentPlayer Player, opponent Player) {
	game.ui.PrintBoard(board)
	move := currentPlayer.GetMove(board, opponent.Symbol(), CHOOSE_CELL)
	board.MakeMove(move, currentPlayer.Symbol())

	winner, symbol := game.rules.Winner(board)

	if winner {
		game.EndGame(board, Winner(symbol))
	} else if board.IsFull() {
		game.EndGame(board, TIE_GAME)
	} else {
		game.Play(board, opponent, currentPlayer)
	}
}

func (game *Game) EndGame(board *Board, message string) {
	game.ui.PrintBoard(board)
	game.ui.DisplayMessage(message)
	game.ui.DisplayMessage(PLAY_AGAIN)
}

func Winner(s string) string {
	return "\n*** Player " + s + " wins! ***\n"
}
