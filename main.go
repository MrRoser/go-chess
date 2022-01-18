package main

import (
	"fmt"
	"os"
	"strconv"
)

var verticalInput = [8]int{
	8, 7, 6, 5, 4, 3, 2, 1,
}

var horizontalInput = [8]string{
	"a", "b", "c", "d", "e", "f", "g", "h",
}

type ChessPiece struct {
	pieceType string
	color     string
}

func main() {
	//initialize values here
	var selectedPiece string

	//game loop
	// for true {

	// }

	//testing area
	gameBoard := create_gameboard()
	draw_gameboard(gameBoard)
	fmt.Printf("\nEnter the position of a piece to move.")
	fmt.Scanf("\nPosition: %s", &selectedPiece)

}

func draw_gameboard(gameboard [8][8]ChessPiece) {

	fmt.Printf("    a   b   c   d   e   f   g   h\n")
	fmt.Printf("  ---------------------------------\n")
	for i := 0; i < 8; i++ {
		rowNumber := 8 - i
		fmt.Printf("%d |", rowNumber)
		for j := 0; j < 8; j++ {
			fmt.Printf(" %s |", gameboard[i][j].pieceType)
		}
		fmt.Printf("\n  ---------------------------------\n")
	}

}

func get_piece(
	selectedPiece string,
	gameboard [8][8]ChessPiece) ChessPiece {

	//break this into multiple functions later

	var verticalPosition int
	var horizontalPosition int

	if len(selectedPiece) != 2 {
		//return error, add later
	}

	for i := 0; i < len(selectedPiece); i++ {
		if i == 0 {
			j := selectedPiece[i]
			k, err := strconv.Atoi(
				string(j))
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			verticalPosition = get_vertical_input_element_valu(k)
		}
		if i == 1 {
			k := string(
				selectedPiece[i])
			horizontalPosition = get_horizontal_input_element_value(k)
		}
	}

	return gameboard[verticalPosition][horizontalPosition]

}

func create_gameboard() [8][8]ChessPiece {

	gameboard := [8][8]ChessPiece{
		{{"C", "black"}, {"H", "black"}, {"B", "black"}, {"Q", "black"},
			{"K", "black"}, {"B", "black"}, {"H", "black"}, {"C", "black"}},

		{{"P", "black"}, {"P", "black"}, {"P", "black"}, {"P", "black"},
			{"P", "black"}, {"P", "black"}, {"P", "black"}, {"P", "black"}},

		{{" ", ""}, {" ", ""}, {" ", ""}, {" ", ""},
			{" ", ""}, {" ", ""}, {" ", ""}, {" ", ""}},

		{{" ", ""}, {" ", ""}, {" ", ""}, {" ", ""},
			{" ", ""}, {" ", ""}, {" ", ""}, {" ", ""}},

		{{" ", ""}, {" ", ""}, {" ", ""}, {" ", ""},
			{" ", ""}, {" ", ""}, {" ", ""}, {" ", ""}},

		{{" ", ""}, {" ", ""}, {" ", ""}, {" ", ""},
			{" ", ""}, {" ", ""}, {" ", ""}, {" ", ""}},

		{{"P", "white"}, {"P", "white"}, {"P", "white"}, {"P", "white"},
			{"P", "white"}, {"P", "white"}, {"P", "white"}, {"P", "white"}},

		{{"C", "white"}, {"H", "white"}, {"B", "white"}, {"K", "white"},
			{"Q", "white"}, {"B", "white"}, {"H", "white"}, {"C", "white"}},
	}

	return gameboard

}

//get better name
func get_vertical_input_element_valu(elementValue int) int {
	//clean this up, no need for a variable declaration before the loop
	var value int
	//add error handling incase its a bad user input
	for i := 0; i < len(verticalInput); i++ {
		if verticalInput[i] == elementValue {
			value = i
		}
	}

	return value
}

//get better name
func get_horizontal_input_element_value(elementValue string) int {
	//clean this up, no need for a variable declaration before the loop
	var value int
	//add error handling incase its a bad user input
	for i := 0; i < len(horizontalInput); i++ {
		if horizontalInput[i] == elementValue {
			value = i
		}
	}

	return value
}
