package game

import "testing"

// Vérifie qu'une nouvelle partie est bien initialisée
func TestNew(t *testing.T) {
	g := New()
	if g.Current != 1 {
		t.Errorf("le premier joueur devrait être 1 (X), got %d", g.Current)
	}
	if g.Over {
		t.Error("la partie ne devrait pas être terminée")
	}
	if g.Winner != 0 {
		t.Errorf("il ne devrait pas y avoir de gagnant, got %d", g.Winner)
	}
}

// Vérifie qu'un coup valide est bien joué
func TestPlay_ValidMove(t *testing.T) {
	g := New()
	err := g.Play(0, 0)
	if err != nil {
		t.Errorf("coup valide refusé : %s", err)
	}
	if g.Board[0][0] != 1 {
		t.Errorf("la case devrait être 1 (X), got %d", g.Board[0][0])
	}
}

// Vérifie que le tour alterne bien entre X (1) et O (2)
func TestPlay_SwitchPlayer(t *testing.T) {
	g := New()
	g.Play(0, 0) // X joue
	if g.Current != 2 {
		t.Errorf("c'est le tour de O (2), got %d", g.Current)
	}
	g.Play(1, 0) // O joue
	if g.Current != 1 {
		t.Errorf("c'est le tour de X (1), got %d", g.Current)
	}
}

// Vérifie qu'on ne peut pas jouer sur une case déjà prise
func TestPlay_CellAlreadyTaken(t *testing.T) {
	g := New()
	g.Play(0, 0)
	err := g.Play(0, 0)
	if err == nil {
		t.Error("aurait dû retourner une erreur pour case déjà prise")
	}
}

// Vérifie qu'on ne peut pas jouer hors du plateau
func TestPlay_OutOfBounds(t *testing.T) {
	g := New()
	err := g.Play(3, 0)
	if err == nil {
		t.Error("aurait dû retourner une erreur pour case hors limites")
	}
	err = g.Play(-1, 0)
	if err == nil {
		t.Error("aurait dû retourner une erreur pour index négatif")
	}
}

// Vérifie la victoire sur une ligne horizontale
func TestWinner_Row(t *testing.T) {
	g := New()
	g.Play(0, 0) // X
	g.Play(1, 0) // O
	g.Play(0, 1) // X
	g.Play(1, 1) // O
	g.Play(0, 2) // X gagne ligne 0
	if g.Winner != 1 {
		t.Errorf("X (1) devrait gagner, got %d", g.Winner)
	}
	if !g.Over {
		t.Error("la partie devrait être terminée")
	}
}

// Vérifie la victoire sur une colonne
func TestWinner_Col(t *testing.T) {
	g := New()
	g.Play(0, 0) // X
	g.Play(0, 1) // O
	g.Play(1, 0) // X
	g.Play(1, 1) // O
	g.Play(2, 0) // X gagne colonne 0
	if g.Winner != 1 {
		t.Errorf("X (1) devrait gagner, got %d", g.Winner)
	}
}

// Vérifie la victoire sur la première diagonale
func TestWinner_Diagonal1(t *testing.T) {
	g := New()
	g.Play(0, 0) // X
	g.Play(0, 1) // O
	g.Play(1, 1) // X
	g.Play(0, 2) // O
	g.Play(2, 2) // X gagne diagonale principale
	if g.Winner != 1 {
		t.Errorf("X (1) devrait gagner, got %d", g.Winner)
	}
}

// Vérifie la victoire sur la deuxième diagonale
func TestWinner_Diagonal2(t *testing.T) {
	g := New()
	g.Play(0, 2) // X
	g.Play(0, 0) // O
	g.Play(1, 1) // X
	g.Play(1, 0) // O
	g.Play(2, 0) // X gagne diagonale inverse
	if g.Winner != 1 {
		t.Errorf("X (1) devrait gagner, got %d", g.Winner)
	}
}

// Vérifie la victoire de O
func TestWinner_PlayerO(t *testing.T) {
	g := New()
	g.Play(2, 2) // X
	g.Play(0, 0) // O
	g.Play(2, 0) // X
	g.Play(1, 1) // O
	g.Play(0, 2) // X
	g.Play(2, 1) // O gagne colonne 1... non, ligne 2 non
	// O gagne diagonale 0,0 -> 1,1 -> 2,2
	g2 := New()
	g2.Play(0, 1) // X
	g2.Play(0, 0) // O
	g2.Play(0, 2) // X
	g2.Play(1, 1) // O
	g2.Play(2, 1) // X
	g2.Play(2, 2) // O gagne diagonale
	if g2.Winner != 2 {
		t.Errorf("O (2) devrait gagner, got %d", g2.Winner)
	}
}

// Vérifie le match nul
func TestDraw(t *testing.T) {
	g := New()
	// X O X
	// X O O
	// O X X
	g.Play(0, 0) // X
	g.Play(0, 1) // O
	g.Play(0, 2) // X
	g.Play(1, 1) // O
	g.Play(1, 0) // X
	g.Play(1, 2) // O
	g.Play(2, 1) // X
	g.Play(2, 0) // O
	g.Play(2, 2) // X
	if !g.Over {
		t.Error("la partie devrait être terminée")
	}
	if g.Winner != 0 {
		t.Errorf("il ne devrait pas y avoir de gagnant, got %d", g.Winner)
	}
}

// Vérifie qu'on ne peut pas jouer après la fin de la partie
func TestPlay_GameOver(t *testing.T) {
	g := New()
	g.Play(0, 0) // X
	g.Play(1, 0) // O
	g.Play(0, 1) // X
	g.Play(1, 1) // O
	g.Play(0, 2) // X gagne
	err := g.Play(2, 2)
	if err == nil {
		t.Error("aurait dû refuser de jouer après la fin de la partie")
	}
}