import { useState } from 'react'
// src/App.jsx
// src/App.jsx
import { useGame } from "./hooks/useGame"
import { Board } from "./components/Board"

export default function App() {
  const { game, error, createGame, playMove } = useGame()

  return (
    <div style={{ display: "flex", flexDirection: "column", alignItems: "center", gap: 35, marginTop: 50 }}>
      <h1>Morpion</h1>

      {!game ? (
        <button onClick={createGame}>Nouvelle partie</button>
      ) : (
        <>
          {game.Over ? (
            <p>{game.Winner ? `${game.Winner} a gagné !` : "Match nul !"}</p>
          ) : (
            <p>Tour du joueur : <strong>{game.Current}</strong></p>
          )}

          {game.Board && <Board board={game.Board} onPlay={playMove} />}

          {error && <p style={{ color: "red" }}>{error}</p>}

          {game.Over && (
            <button onClick={createGame}>Rejouer</button>
          )}
        </>
      )}
    </div>
  )
}
