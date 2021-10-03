/*
My app "tic tac toe"

*/
package main

import (
	"fmt"
	"strings"
)

// stores the state of the fields and counter of free fields
type GameGrid struct {

	// game filed 3x3 stored as array
	arrFields [9]int8

	// changes when the player placed on the field (makes a move)
	freeFields uint8

	// changes when the player is win
	winner int8

	// player map for Print
	mapPlayer map[int8]string
}

// Reset Game fields to initial values
func (grid *GameGrid) ResetGameGrid() {
	grid.arrFields = [9]int8{}
	grid.freeFields = 9
	grid.winner = 0
	grid.mapPlayer = map[int8]string{-1: "O", 0: "_", 1: "X"}
}

// Print game fields
func (grid GameGrid) PrintFields() {

	// print grid
	for idx, val := range grid.arrFields {
		remainder := (idx + 1) % 3
		newLine := ""

		if remainder == 0 {
			newLine = "\n"
		} else {
			newLine = ""
		}

		fmt.Printf("%v "+newLine, grid.mapPlayer[val])
	}
}

/*
	Place player to spicific position on grid (3x3).
	Return status of operation and set counter of empty(zero) fields.
	Check whos is win.

	incoming params:
		player - type of player "-1"="O", "1"="X"
		row - row number 1, 2 or 3
		col - column number 1, 2 or 3

	return params:
		bool - "true" player placed to field, "false" field is taken

*/
func (grid *GameGrid) placePlayer(player int8, row uint8, col uint8) bool {

	// on success return true
	result := true

	// adjust col and row values
	row -= 1
	col -= 1

	// check row and col values
	if (row < 0) || (row > 2) {
		result = false
	}

	if (col < 0) || (col > 2) {
		result = false
	}

	// set player position
	if result {
		var position uint8 = row*3 + col
		if grid.arrFields[position] == 0 {
			grid.arrFields[position] = player
			grid.freeFields -= 1
		} else {
			result = false
		}
	}

	// checking Winner
	grid.winner = grid.CheckWinner()

	return result
}

/*
Test all possible winning combinations and return who is win ('X', 'O' or no body)
*/
func (grid GameGrid) CheckWinner() int8 {

	/*
		List of tests to determine the winner:
			3 test to check rows
			3 test to check collumns
			2 test to check diagonal (cross)
	*/

	var winner int8 = 0
	var fields [3]int8

	// test rows
	for i := 0; i < 3; i++ {
		fields = [3]int8{grid.arrFields[0+i*3], grid.arrFields[1+i*3], grid.arrFields[2+i*3]}
		winner = lookingForWinner(fields)
		if winner != 0 {
			return winner
		}
	}

	// test cols
	for i := 0; i < 3; i++ {
		fields = [3]int8{grid.arrFields[0+i], grid.arrFields[3+i], grid.arrFields[6+i]}
		winner = lookingForWinner(fields)
		if winner != 0 {
			return winner
		}
	}

	// test diagonal [0,4,8] and [2,4,6] -> 0,8 and 2,6 => delta is 2
	var d uint8 = 0
	for d < 3 {
		fields = [3]int8{grid.arrFields[0+d], grid.arrFields[4], grid.arrFields[8-d]}
		winner = lookingForWinner(fields)
		if winner != 0 {
			return winner
		}
		d += 2
	}

	// no winner
	return winner
}

/*
Checking whos is win.
*/
func lookingForWinner(fields3 [3]int8) int8 {
	var sum int8 = 0
	var mul int8 = 1

	for _, val := range fields3 {
		sum += val
		mul *= val
	}

	if mul == 0 {
		return 0
	}

	switch sum {
	case 3:
		return 1
	case -3:
		return -1
	default:
		return 0
	}

}

// Game object
type Game struct {

	// Number of rounds
	round uint

	// Score
	scoreX    uint
	scoreO    uint
	scoreDraw uint

	// Game grid with states
	grid GameGrid

	// Current player
	player int8

	// Game state
	isPlay bool
}

// Initialising game variables
func (g *Game) NewGame() {
	g.round, g.scoreDraw, g.scoreO, g.scoreX = 0, 0, 0, 0
	g.isPlay = true

	fmt.Println("New game started!")
}

// launch new round
func (g *Game) NewRound() {

	g.player = 1
	g.round += 1

	g.grid.ResetGameGrid()

	fmt.Printf("Starting new round: %v\n", g.round)
	g.PrintScore()
}

// Check end of round. Return 'true' if round is over and set score variables.
func (g *Game) EndOfRound() bool {

	if (g.grid.freeFields == 0) || (g.grid.winner != 0) {

		// set score
		switch g.grid.winner {
		case -1:
			g.scoreO += 1
		case 1:
			g.scoreX += 1
		default:
			g.scoreDraw += 1
		}

		// round is over
		return true
	}

	// continue
	return false
}

// Print Score
func (g Game) PrintScore() {
	fmt.Printf("Score X: %v, O: %v, Draw: %v\n", g.scoreX, g.scoreO, g.scoreDraw)
}

func main() {

	var row, col uint8
	var answer string
	var game Game

	game.NewGame()

	// play a game
	for game.isPlay {

		game.NewRound()

		// play a game round
		for !game.EndOfRound() {
			game.grid.PrintFields()

			// player inpunt from console
			fmt.Printf("Player '%v' enter row and col: ", game.grid.mapPlayer[game.player])
			fmt.Scanln(&row, &col)

			if game.grid.placePlayer(game.player, row, col) {
				// change player
				game.player *= -1
			} else {
				fmt.Println("Plaese select another field")
			}
		}

		game.grid.PrintFields()
		fmt.Printf("Round %v is over\n", game.round)

		// congratulations to the winner
		if game.grid.winner != 0 {
			fmt.Printf("Player '%v' won!\n", game.grid.mapPlayer[game.grid.winner])
		} else {
			fmt.Println("Friendship won!")
		}

		game.PrintScore()

		fmt.Println("Do you want to play another round (yes/no)?")
		fmt.Scanln(&answer)
		answer = strings.ToLower(answer)

		switch answer {
		case "y", "yes":
			game.isPlay = true
		default:
			game.isPlay = false
		}
	}

}
