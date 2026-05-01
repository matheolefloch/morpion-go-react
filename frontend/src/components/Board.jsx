// src/components/Board.jsx
import { Cell } from "./Cell"

export function Board({ board, onPlay }) {
  return (
    <div style={{ display: "grid", gridTemplateColumns: "repeat(3, 100px)", gap: 4 }}>
      {board.map((row, r) =>
        row.map((cell, c) => (
          <Cell key={`${r}-${c}`} value={cell} onClick={() => onPlay(r, c)} />
        ))
      )}
    </div>
  )
}
