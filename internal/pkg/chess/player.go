package chess

// Color represents the two teams of chess
type Color string

const (
	White Color = "white"
	Black Color = "black"
)

// Player represents a chess player. A chess player is either the white or black player, and has a piece. This could be
// replaced with some set of pieces for a full game.
type Player struct {
	Color Color
	Piece *Piece
}

// Move performs a move as this player. In this simplified game, the player's single piece is moved.
// In a full game, logic to select a piece to move would go here.
func (p *Player) Move() {
	p.Piece.Move()
}

// CheckForWin checks if we've hit a victory condition. In this case, if our piece is able to make a move that would capture
// the opponent's piece.
func (p *Player) CheckForWin(opponent *Player) bool {
	return p.Piece.HasIntercept(opponent.Piece)
}
