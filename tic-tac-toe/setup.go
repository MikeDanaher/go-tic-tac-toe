package tic_tac_toe

import "os"

type Setup struct {
	config Config
}

func NewSetup(config Config) *Setup {
	return &Setup{config}
}

func (s *Setup) Run() {
	rules := NewRules(s.config.WinningLines())
	board := NewBoard(s.config.BoardSize())
	ui := s.config.UI()
	game := NewGame(ui, rules)

	p1, p2 := choosePlayers(ui, rules)
	chooseSymbols(ui, p1, p2)
	firstPlayer := chooseFirstPlayer(ui)

	if firstPlayer == 1 {
		game.Play(board, p1, p2)
	} else {
		game.Play(board, p2, p1)
	}
}

func choosePlayers(ui UI, rules *Rules) (player1 Player, player2 Player) {
	ui.DisplayMessage(MAIN_MENU)
	choice := ui.GetNumber()

	switch choice {
	case 1:
		player1 = NewHumanPlayer(ui, rules)
		player2 = NewHumanPlayer(ui, rules)
	case 2:
		player1 = NewHumanPlayer(ui, rules)
		player2 = NewHumanPlayer(ui, rules)
	case 3:
		os.Exit(0)
	default:
		return choosePlayers(ui, rules)
	}
	return
}

func chooseSymbols(ui UI, p1 Player, p2 Player) {
	ui.DisplayMessage(CHOOSE_YOUR_SYMBOL)
	s1 := ui.GetString()
	p1.SetSymbol(s1)

	ui.DisplayMessage(CHOOSE_OPPONENTS_SYMBOL)
	s2 := ui.GetString()
	p2.SetSymbol(s2)
}

func chooseFirstPlayer(ui UI) int {
	ui.DisplayMessage(CHOOSE_FIRST_PLAYER)
	first := ui.GetNumber()

	switch first {
	case 1:
		return 1
	case 2:
		return 2
	}
	return chooseFirstPlayer(ui)
}
