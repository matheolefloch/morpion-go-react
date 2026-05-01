package game

import "errors"

type Game struct {
	Board   [3][3]int `json:"Board"`   // 0 : empty, 1 : X, 2 : O
	Current int       `json:"Current"` // X : 1, O : 2
	Winner  int       `json:"Winner"`
	Over    bool      `json:"Over"`
}

func New() *Game {
	return &Game{Current: 1}
}

func (g *Game) Play(row, col int) error {
	if g.Over {
		return errors.New("la partie est terminée")
	}
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return errors.New("case invalide")
	}
	if g.Board[row][col] != 0 {
		return errors.New("case déjà prise")
	}

	g.Board[row][col] = g.Current
	g.checkWinner()

	if !g.Over {
		g.switchPlayer()
	}
	return nil
}

func (g *Game) switchPlayer() {
	g.Current = 3 - g.Current
}

func (g *Game) checkWinner() {
	winner := g.getWinner()
	if winner != 0 {
		g.Winner = winner
		g.Over = true
		return
	}

	full := true
	for _, row := range g.Board {
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
		g.Over = true
	}
}

func (g *Game) getWinner() int {
	b := g.Board

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
