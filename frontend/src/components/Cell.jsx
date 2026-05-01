// src/components/Cell.jsx
export function Cell({ value, onClick }) {
    const isX = value === 1
    const isO = value === 2
  return (
    <button
      onClick={onClick}
      style={{
        width: 100,
        height: 100,
        cursor: value ? "default" : "pointer",
        backgroundColor: "#fff",
        border: "2px solid #ddd",
        borderRadius: 12,
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        boxShadow: "0 2px 6px rgba(0,0,0,0.08)",
      }}
    >
      <span
        style={{
          fontSize: isO ? 105 : 56,   // 👈 O plus grand que X
          fontWeight: "bold",
          color: isX ? "#e74c3c" : "#3498db",
          transform: "translateY(-5px)",
        }}
      >
        {isX ? "✕" : isO ? "○" : ""}
      </span>
    </button>
  )
}
