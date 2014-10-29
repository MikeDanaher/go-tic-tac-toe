package tic_tac_toe

import "os"

type ConsoleSetup struct {
	board   *Board
	game    *Game
	Player1 Player
	Player2 Player
}

func NewConsoleSetup(config Config) *ConsoleSetup {
	ui := NewConsoleUI(config.Input(), config.Output())
	gameType := chooseGame(ui)

	rules := NewRules(config.WinningLines())
	board := NewBoard(config.BoardSize())
	game := NewGame(ui, rules)
	player1, player2 := buildPlayers(ui, rules, gameType)

	ui.DisplayMessage(DESCRIBE_GAME)
	ui.DisplayMessage(CELL_NUMBERS)

	return &ConsoleSetup{board, game, player1, player2}
}

func (s *ConsoleSetup) Run() {
	s.game.Play(s.board, s.Player1, s.Player2)
}

func chooseGame(ui UI) string {
	ui.DisplayMessage(MAIN_MENU)
	choice := ui.GetNumber()

	switch choice {
	case 1:
		return HUMAN_v_HUMAN
	case 2:
		return HUMAN_v_COMPUTER
	case 3:
		return COMPUTER_v_COMPUTER
	case 4:
		os.Exit(0)
	}
	return chooseGame(ui)
}

func buildPlayers(ui UI, rules *Rules, gameType string) (Player, Player) {
	var p1, p2 Player
	switch gameType {
	case HUMAN_v_HUMAN:
		p1 = NewHumanPlayer(ui, rules)
		p2 = NewHumanPlayer(ui, rules)
	case HUMAN_v_COMPUTER:
		p1 = NewHumanPlayer(ui, rules)
		p2 = NewComputerPlayer(rules)
	case COMPUTER_v_COMPUTER:
		p1 = NewComputerPlayer(rules)
		p2 = NewComputerPlayer(rules)
	}
	chooseSymbols(ui, p1, p2)
	return chooseFirstPlayer(ui, p1, p2)
}

func chooseSymbols(ui UI, p1 Player, p2 Player) {
	ui.DisplayMessage(CHOOSE_YOUR_SYMBOL)
	s1 := ui.GetString()
	p1.SetSymbol(s1)

	ui.DisplayMessage(CHOOSE_OPPONENTS_SYMBOL)
	s2 := ui.GetString()
	p2.SetSymbol(s2)
}

func chooseFirstPlayer(ui UI, p1 Player, p2 Player) (Player, Player) {
	ui.DisplayMessage(CHOOSE_FIRST_PLAYER)
	first := ui.GetNumber()

	switch first {
	case 1:
		return p1, p2
	case 2:
		return p2, p1
	}
	return chooseFirstPlayer(ui, p1, p2)
}
