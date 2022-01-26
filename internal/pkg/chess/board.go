package chess

import (
	"fmt"
	"time"
)

// Board represents a chess board. A board has two players, which have pieces, which have positions and types.
type Board struct {
	// Given that there can only ever be two chess players, defined fields seemed to make more sense than a data structure
	// like a map or a list.
	WhitePlayer *Player
	BlackPlayer *Player
}

// InitBoard initializes a chess board. In this case, a simplified game of chess with just two pieces.
func InitBoard() Board {
	return Board{
		WhitePlayer: &Player{
			Color: White,
			Piece: NewBishop(White, Bishop, 2, 2),
		},
		BlackPlayer: &Player{
			Color: Black,
			Piece: NewRook(Black, Rook, 7, 0),
		},
	}
}

// String prints the current board state.
func (b *Board) String() string {
	str := "----Current Board Positions----\n"
	str += fmt.Sprintf("%v\n", b.WhitePlayer.Piece)
	str += fmt.Sprintf("%v\n", b.BlackPlayer.Piece)

	return str
}

// Run performs up to 15 chess moves, stopping early if there's a winner.
func (b *Board) Run() {
	fmt.Println("----Starting a game of chess----")
	fmt.Printf("%s", b)

	// run for up to 15 turns. Assuming that a turn consists of a move by both white and black
	for turn := 0; turn < 15; turn++ {
		fmt.Printf("\n----Running turn %d----\n", turn)

		b.WhitePlayer.Move()

		// check if black can capture white after white's move, as it's black's turn to make a move
		if b.BlackPlayer.CheckForWin(b.WhitePlayer) {
			fmt.Printf("%s", b)
			fmt.Println("The black rook is able to capture the white bishop")
			fmt.Println("Black is the winner!")
			return
		}

		b.BlackPlayer.Move()

		// check if white can capture black after black's move, as it's white's turn to make a move
		if b.WhitePlayer.CheckForWin(b.BlackPlayer) {
			fmt.Printf("%s", b)
			fmt.Println("The white bishop is able to capture the black rook")
			fmt.Println("White is the winner!")
			return
		}

		fmt.Printf("%s", b)

		// sleep for the sake of following terminal output
		time.Sleep(2 * time.Second)
	}

	fmt.Println("No winner was found after 15 moves. Stalemate!")
}
