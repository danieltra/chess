package chess

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/adam-lavrik/go-imath/ix"
)

type PieceType string

const (
	Rook   PieceType = "rook"
	Bishop PieceType = "bishop"
)

// seed the rand package once when the package is loaded
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Mover enables unique movement implementations for different types of chess pieces. When a Piece is initialized,
// a type specific mover is provided. In this game, implementations are provided for the Rook and Bishop. However,
// for the scope of this game the bishop's move is a no-op. For a full game, the bishop and all remaining pieces could
// be implemented.
type Mover interface {
	Move(position *Position)
	HasIntercept(me *Piece, opponent *Piece) bool
}

/*****************************************
* General chess piece logic
*****************************************/

// Piece represents a piece on the chess board. A piece as a position, and a specific type of movement that it's capable of.
type Piece struct {
	Color    Color
	Type     PieceType
	Position *Position
	Mover
}

// String enables pretty printing of the piece
func (p *Piece) String() string {
	return fmt.Sprintf("%s %s: %s", p.Color, p.Type, p.Position)
}

// Move calls the corresponding Move() function on the configured Mover(0)
func (p *Piece) Move() {
	p.Mover.Move(p.Position)
}

// HasIntercept call the correpsonding has intercept function on the configured mover
func (p *Piece) HasIntercept(opponent *Piece) bool {
	return p.Mover.HasIntercept(p, opponent)
}

/*****************************************
* Logic related to the bishop
*****************************************/

// NewBishop initializes a bishop in the specified position.
func NewBishop(color Color, pieceType PieceType, file, rank int) *Piece {
	return &Piece{
		Color:    color,
		Type:     pieceType,
		Position: newPosition(file, rank),
		Mover:    &BishopMover{},
	}
}

type BishopMover struct{}

// The move operation for the bishop type is a no-op in this exercise, as the bishop is stationary
func (b *BishopMover) Move(position *Position) {}

// HasIntercept checks if a bishop can intercept a piece.
// A bishop can intercept a piece if the absolute value of the difference between the rank and file of both pieces are equal.
// For example, if the bishop is at b2 and the rook is at d4 -- the bishop has an intercept. (d - b) == 2 and (4 - 2) == 2.
// This is an artifact of the fact that a diagonal move in chess requires changing the rank and file by the same magnitude.
func (b *BishopMover) HasIntercept(me *Piece, opponent *Piece) bool {
	return ix.Abs(me.Position.file-opponent.Position.file) ==
		ix.Abs(me.Position.rank-opponent.Position.rank)
}

/*****************************************
* Logic related to the rook
*****************************************/

// NewRook initializes a rook in the specified position
func NewRook(color Color, pieceType PieceType, file, rank int) *Piece {
	return &Piece{
		Color:    color,
		Type:     pieceType,
		Position: newPosition(file, rank),
		Mover:    &RookMover{},
	}
}

type RookMover struct{}

// Move enables a rook to move up, or to the right. Two six sided dice are rolled to determine the number of spaces of the movement.
// A coin flip is performed to determine the direction.
func (r *RookMover) Move(position *Position) {
	fmt.Println("Calculating the rook's move")
	// generate a random value between 0 and 1. On 0 (tails) shift the File, on 1 (heads) shift the rank.
	// The value moved is the sum of two six sized dice.
	shifters := map[int]func(int){
		0: position.shiftFile,
		1: position.shiftRank,
	}

	d2ToString := map[int]string{
		0: "file",
		1: "rank",
	}

	// a coin toss is equal to a 2 sided die
	direction := d2()
	fmt.Printf("The outcome of the coin toss was %d\n", direction)

	// rolling two six sided die, per the requirement. 1 12 sided die wouldn't have the same distribution of outcomes.
	die1 := d6()
	die2 := d6()
	degree := die1 + die2
	fmt.Printf("The die rolls were %d and %d for a sum of %d\n", die1, die2, degree)

	fmt.Printf("changing the rook's %s by %d spaces\n", d2ToString[direction], degree)

	shifters[direction](degree)
}

// HasIntercept checks if a rook can intercept a piece.
// A rook can intercept a piece if that piece is in either the same rank or file.
func (r *RookMover) HasIntercept(me *Piece, opponent *Piece) bool {
	if me.Position.file == opponent.Position.file || me.Position.rank == opponent.Position.rank {
		return true
	}

	return false
}

/*****************************************
* Helpers for random number generation
*****************************************/

// d6 simulates a 6 sided die
func d6() int {
	return rand.Intn(6)
}

// d2 simulates a coinflip
func d2() int {
	return rand.Intn(2)
}
