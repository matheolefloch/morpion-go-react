# 🎮 Morpion — Go + React

Un morpion fullstack développé en **Go** (backend) et **React** (frontend), réalisé comme projet d'apprentissage pour découvrir l'écosystème Go.

---

## 🚀 Stack technique

| Côté | Techno |
|------|--------|
| Backend | Go, Gin |
| Frontend | React |
| Communication | API REST (JSON) |

---

## 📁 Structure du projet

```
morpion/
├── go.mod
├── main.go           # Point d'entrée, serveur HTTP
├── game/
│   ├── game.go       # Logique métier du jeu
│   └── game_test.go  # Tests unitaires
└── frontend/         # Application React
```

---

## ⚙️ Lancer le projet

### Prérequis

- [Go 1.22+](https://go.dev/dl)
- [Node.js 18+](https://nodejs.org)

### Backend

```bash
# Cloner le projet
git clone https://github.com/matheolefloch/morpion-go-react.git
cd morpion-go-react

# Installer les dépendances Go
go mod tidy

# Lancer le serveur
go run main.go
# → API disponible sur http://localhost:8080
```

### Frontend

```bash
cd frontend
npm install
npm run dev
# → App disponible sur http://localhost:5173
```

### Tests

```bash
go test ./game/...
```

---

## 🎯 Fonctionnalités

- Partie en local à deux joueurs (X et O)
- Détection automatique de la victoire et du match nul
- Réinitialisation de la partie
- API REST pour interagir avec la logique du jeu

## 🔌 API

| Méthode | Route | Description |
|---------|-------|-------------|
| `POST` | `/game` | Créer une nouvelle partie |
| `GET` | `/game/:id` | Récupérer l'état d'une partie |
| `POST` | `/game/:id/move` | Jouer un coup |

**Exemple — jouer un coup :**

```bash
curl -X POST http://localhost:8080/game/1/move \
  -H "Content-Type: application/json" \
  -d '{"row": 0, "col": 1}'
```

---

## 📖 Ce que j'ai appris

- Structurer un projet Go en packages
- Créer une API REST avec **Gin**
- Gérer les erreurs à la façon Go
- Écrire des **tests unitaires** avec le package `testing`
- Connecter un frontend React à une API

## 🔮 Améliorations possibles

- Multijoueur en temps réel avec **WebSockets** (goroutines + channels)
- Persistance des parties en base de données (PostgreSQL)
- Système de matchmaking / salles de jeu
- IA simple pour jouer contre l'ordinateur

---

## 👤 Auteur

**Ton Nom** — [github.com/ton-pseudo](https://github.com/ton-pseudo)

> Projet réalisé en 2 jours dans le cadre d'un apprentissage de Go pour une candidature de stage.
