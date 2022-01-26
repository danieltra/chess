package main

import (
	"github.com/danieltra/chess/internal/pkg/chess"
)

func main() {
	board := chess.InitBoard()

	board.Run()
}
