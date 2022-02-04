package main

import (
	"errors"
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
	var selectedPosition string

	gameBoard := create_gameboard()

	//game loop
	for {

		draw_gameboard(gameBoard)
		fmt.Printf("\nEnter the position of a piece to move.")
		fmt.Scanf("\nPosition: %s", &selectedPiece)
		startingPosition := get_piece(selectedPiece, gameBoard)
		fmt.Printf("\nEnter the position of where the piece will move.")
		fmt.Scanf("\nPosition: %s", &selectedPosition)
		newPosition := get_piece(selectedPosition, gameBoard)
		updateGameBoard(startingPosition, newPosition)

	}

	// //testing area
	// gameBoard := create_gameboard()
	// draw_gameboard(gameBoard)
	// fmt.Printf("\nEnter the position of a piece to move.")
	// fmt.Scanf("\nPosition: %s", &selectedPiece)
	// startingPosition := get_piece(selectedPiece, gameBoard)
	// fmt.Printf("\nEnter the position of where the piece will move.")
	// fmt.Scanf("\nPosition: %s", &selectedPosition)
	// newPosition := get_piece(selectedPosition, gameBoard)
	// updateGameBoard(startingPosition, newPosition)

	// print(color.Ize(color.Cyan, "Test"))
	// println(color.Ize(color.Green, "Test2"))

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
	var inputError error //change this later on
	var verticalPosition int
	var horizontalPosition int

	if selectedPiece == "exit" {
		println("Exit")
		os.Exit(3)
	}

	// if len(selectedPiece) != 2 {
	// 	//return error, add later
	// }

	for i := 0; i < len(selectedPiece); i++ {
		if i == 0 {
			j := selectedPiece[i]
			k, err := strconv.Atoi(
				string(j))
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			verticalPosition, inputError = get_vertical_input_element_value(k)
			if inputError != nil {
				println(inputError)
				os.Exit(3)
			}
		}
		if i == 1 {
			k := string(
				selectedPiece[i])
			horizontalPosition, inputError = get_horizontal_input_element_value(k)
			if inputError != nil {
				println(inputError)
				os.Exit(3)
			}
		}
	}

	return gameboard[verticalPosition][horizontalPosition]

}

func updateGameBoard(
	startPosition ChessPiece,
	newPosition ChessPiece) {

	newPosition.color = startPosition.color
	newPosition.pieceType = startPosition.pieceType
	startPosition.color = ""
	startPosition.pieceType = " "
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
func get_vertical_input_element_value(
	elementValue int) (
	int,
	error) {
	for i := 0; i < len(verticalInput); i++ {
		if verticalInput[i] == elementValue {
			return i, nil
		}
	}
	return -1, errors.New(
		"input error for vertical position")
}

//get better name
func get_horizontal_input_element_value(
	elementValue string) (
	int,
	error) {
	for i := 0; i < len(horizontalInput); i++ {
		if horizontalInput[i] == elementValue {
			return i, nil
		}
	}
	return -1, errors.New(
		"input error for horizontal position")
}
