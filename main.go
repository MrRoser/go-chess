package main

import "fmt"

type BoardPosition struct {
	gamePiece string
	color     string
}

type ChessPiece struct {
	pieceType string
	color     string
}

func main() {
	//initialize values here

	//game loop
	// for true {

	// }

	//testing area
	draw_gameboard(create_gameboard())
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
