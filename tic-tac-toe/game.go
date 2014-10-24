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
	move := currentPlayer.GetMove(board, CHOOSE_CELL)
	board.MakeMove(move, currentPlayer.Symbol())

	winner, symbol := game.rules.Winner(board)

	if winner {
		game.ui.PrintBoard(board)
		game.ui.DisplayMessage(Winner(symbol))
	} else if board.IsFull() {
		game.ui.PrintBoard(board)
		game.ui.DisplayMessage(TIE_GAME)
	} else {
		game.Play(board, opponent, currentPlayer)
	}
}

func Winner(s string) string {
	return "\nPlayer " + s + " wins!\n"
}
