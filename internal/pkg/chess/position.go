package chess

import "fmt"

var fileIntToString = map[int]string{
	0: "a",
	1: "b",
	2: "c",
	3: "d",
	4: "e",
	5: "f",
	6: "g",
	7: "h",
}

// Represents the position of a chess piece on a board
type Position struct {
	file int
	rank int
}

func (p *Position) String() string {
	// convert the file to a string, and shift the rank by 1 since the underlying integer value can be between 0 and 7.
	return fmt.Sprintf("%s%v", fileIntToString[p.file], (p.rank + 1))
}

// initializes a position with the provided rank and file
func newPosition(file int, rank int) *Position {
	return &Position{
		file,
		rank,
	}
}

// Shifts left or right a number of files. A negative number will shift the position to the left,
// a positive number will shift the position to the right. This function auto-wraps when moving off the board.
func (p *Position) shiftFile(degree int) {
	p.file = shift(p.file, degree)
}

// Shifts up or down a number of ranks. A negative number will shift the position down,
// a positive number will shift the position up. This function auto-wraps when moving off the board.
func (p *Position) shiftRank(degree int) {
	p.rank = shift(p.rank, degree)
}

// A helper that enables shifting an integer by some amount, wrapping if the value leaves the bound of 0 to 7.
// This is purpose-specific to representing the position of a piece on a chess board.
func shift(currentPosition, degree int) int {
	currentPosition = (currentPosition + degree) % 8
	// if we've gone negative, wrap back. For example, at position -1 we'll wrap to 7 (8 + -1).
	if currentPosition < 0 {
		currentPosition = 8 + currentPosition
	}

	return currentPosition
}
