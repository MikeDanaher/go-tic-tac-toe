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
	move := currentPlayer.GetMove(board, ChooseCell)
	board.MakeMove(move, currentPlayer.Symbol())

	winner, symbol := game.rules.Winner(board)

	if winner {
		game.ui.PrintBoard(board)
		game.ui.DisplayMessage("Player " + symbol + " wins!")
		game.ui.DisplayMessage(MainMenu)
	} else if board.IsFull() {
		game.ui.DisplayMessage("It's a tie!")
	} else {
		game.Play(board, opponent, currentPlayer)
	}
}
