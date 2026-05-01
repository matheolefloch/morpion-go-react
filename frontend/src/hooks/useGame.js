// src/hooks/useGame.js
import { useState } from "react"

const API = "http://localhost:8080"

export function useGame() {
  const [game, setGame] = useState(null)
  const [gameId, setGameId] = useState(null)
  const [error, setError] = useState(null)

  async function createGame() {
    const res = await fetch(`${API}/game`, { method: "POST" })
    const data = await res.json()
    setGameId(data.id)

    const res2 = await fetch(`${API}/game/${data.id}`)
    setGame(await res2.json())
  }

  async function playMove(row, col) {
    if (!gameId) return
    const res = await fetch(`${API}/game/${gameId}/move`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ row, col }),
    })
    const data = await res.json()
    console.log("Move reçu :", data)
    if (data.error) {
      setError(data.error)
      return
    }
    setError(null)
    setGame(data)
  }

  return { game, error, createGame, playMove }
}
