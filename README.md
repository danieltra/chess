# chess
This repo contains a simulation of a simplified game of chess, with only two pieces. When the program is run the simulation is run until the game is complete, up to 15 turns. White wins if it has a move to capture black, black wins if has a move to capture white.

## Running the Application

The app can be run via `go run cmd/main.go`.

## Assumptions
- After white's move, I check if black has a winning move. After black's move, I check if white has a winning move. The assumption here is that the opponent would make the winning move on their turn if they could. A traditional checkmate wouldn't be possible with this board setup.
- There's currently no detection on if black would move through white on it's turn, as I don't believe that's a possibility. This would require black to already be in a position to capture white (and would have therefore already won), and since white isn't moving it would never move itself into that position.
- I also don't believe it's possible for black to directly land on white for the same reason.
- I assumed there will only ever be two players
- I assumed there would only ever be 1 chess piece per player
- I treated a coin toss like a 2 sided die (a value of 0 or 1)
- I assumed a game would only ever be 15 turns, as that was specified in the challenge
