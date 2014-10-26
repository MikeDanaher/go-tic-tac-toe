package tic_tac_toe

import "os"

type ConsoleSetup struct {
	board   *Board
	game    *Game
	player1 Player
	player2 Player
}

const (
	MAIN_MENU = "\nTIC-TAC-TOE in GO\n" +
		"=====================================================\n" +
		"1. Human vs. Human game\n" +
		"2. Human vs. Computer game\n" +
		"3. Exit\n" +
		"Play the game by entering the number of your choice: "
	CHOOSE_CELL             = "Please select an available cell [1-9]: "
	INVALID_NUMBER          = "Please enter a valid number: "
	INVALID_MOVE            = "Invalid move, try again: "
	CHOOSE_YOUR_SYMBOL      = "Enter your symbol: "
	CHOOSE_OPPONENTS_SYMBOL = "Enter your opponent's symbol: "
	CHOOSE_FIRST_PLAYER     = "\n1. You\n" +
		"2. Opponent\n" +
		"Who will go first? "
	TIE_GAME   = "\n*** It's a Tie! ***\n"
	PLAY_AGAIN = "\nPlay Again?\n"
	HUMAN      = "HUMAN"
	COMPUTER   = "COMPUTER"
)

func NewConsoleSetup(config Config) *ConsoleSetup {
	ui := NewConsoleUI(os.Stdin, os.Stdout)
	gameType := chooseGame(ui)

	rules := NewRules(config.WinningLines())
	board := NewBoard(config.BoardSize())
	game := NewGame(ui, rules)
	player1, player2 := buildPlayers(ui, rules, gameType)

	return &ConsoleSetup{board, game, player1, player2}
}

func (s *ConsoleSetup) Run() {
	s.game.Play(s.board, s.player1, s.player2)
}

func chooseGame(ui UI) string {
	ui.DisplayMessage(MAIN_MENU)
	choice := ui.GetNumber()

	switch choice {
	case 1:
		return HUMAN
	case 2:
		return COMPUTER
	case 3:
		os.Exit(0)
	}
	return chooseGame(ui)
}

func buildPlayers(ui UI, rules *Rules, gameType string) (Player, Player) {
	var p1, p2 Player
	switch gameType {
	case HUMAN:
		p1 = NewHumanPlayer(ui, rules)
		p2 = NewHumanPlayer(ui, rules)
	case COMPUTER:
		p1 = NewHumanPlayer(ui, rules)
		p2 = NewComputerPlayer()
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
