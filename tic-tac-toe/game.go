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
		game.End(board, PrintWinner(symbol))
	} else if board.IsFull() {
		game.End(board, TIE_GAME)
	} else {
		game.Play(board, opponent, currentPlayer)
	}
}

func (game *Game) End(board *Board, message string) {
	game.ui.DisplayMessage(message)
	game.ui.PrintBoard(board)
	game.ui.DisplayMessage(PLAY_AGAIN)
}
