package game

import "errors"

type Game struct {
	board   [3][3]int // 0 : empty, 1 : X, 2 : O
	current int       // X : 1, O : 2
	winner  int
	over    bool
}

func New() *Game {
	return &Game{current: 1}
}

func (g *Game) Play(row, col int) error {
	if g.over {
		return errors.New("la partie est terminée")
	}
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return errors.New("case invalide")
	}
	if g.board[row][col] != 0 {
		return errors.New("case déjà prise")
	}

	g.board[row][col] = g.current
	g.checkWinner()

	if !g.over {
		g.switchPlayer()
	}
	return nil
}

func (g *Game) switchPlayer() {
	g.current = 3 - g.current
}

func (g *Game) checkWinner() {
	winner := g.getWinner()
	if winner != 0 {
		g.winner = winner
		g.over = true
		return
	}

	full := true
	for _, row := range g.board {
		for _, cell := range row {
			if cell == 0 {
				full = false
			}
		}
		if !full {
			break
		}
	}

	if full {
		g.over = true
	}
}

func (g *Game) getWinner() int {
	b := g.board

	for i := 0; i < 3; i++ {
		// Lines
		if b[i][0] != 0 && b[i][0] == b[i][1] && b[i][1] == b[i][2] {
			return b[i][0]
		}
		// Columns
		if b[0][i] != 0 && b[0][i] == b[1][i] && b[1][i] == b[2][i] {
			return b[0][i]
		}
	}

	// First diagonal
	if b[0][0] != 0 && b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		return b[0][0]
	}

	// Second diagonal
	if b[0][2] != 0 && b[0][2] == b[1][1] && b[1][1] == b[2][0] {
		return b[0][2]
	}

	return 0
}
