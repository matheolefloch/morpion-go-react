package main

import (
	"fmt"
	"morpion/game"
	"net/http"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Save the game stats
var (
	games   = map[string]*game.Game{}
	mu      sync.Mutex
	counter = 0
)

func main() {
	// Use of gin Framework
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/game", createGame)
	r.GET("/game/:id", getGame)
	r.POST("/game/:id/move", playMove)

	r.Run(":8080")
}

// POST /game : Create a new game
func createGame(c *gin.Context) {
	mu.Lock()
	counter++
	id := fmt.Sprintf("%d", counter)
	games[id] = game.New()
	mu.Unlock()

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GET /game/:id : Get the state of the game
func getGame(c *gin.Context) {
	id := c.Param("id")

	mu.Lock()
	g, ok := games[id]
	mu.Unlock()

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Partie introuvable"})
		return
	}

	c.JSON(http.StatusOK, *g)
}

// POST /game/:id/move : Play a move
func playMove(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Row int `json:"row"`
		Col int `json:"col"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "body invalide"})
		return
	}

	mu.Lock()
	g, ok := games[id]
	mu.Unlock()

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "partie introuvable"})
		return
	}

	if err := g.Play(body.Row, body.Col); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, *g)
}
