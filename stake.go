package main

import(
	"fmt"
	"math/rand"
	"time"
	"strings"
)
const SIZE = 5
var board [SIZE][SIZE]rune    
var displayBoard [SIZE][SIZE]rune 

func initializeBoard(){
	for i:=0; i<SIZE;i++{
		for j:=0; j<SIZE; j++{
			board[i][j]=' '

		}
	}
}

func initializeDisplayBoard(){
	for i:=0;i<SIZE;i++{
		for j:=0;j<SIZE;j++{
			displayBoard[i][j]='?'
		}
	}
}

func RandomSpotGenerator(){
	source:= rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	num1 := rng.Intn(SIZE)
	num2 := rng.Intn(SIZE)
	board[num1][num2] = 'X'


}


func printSeparator() {
	fmt.Println(strings.Repeat("-", SIZE*4))
}

func printtheBoard(board [SIZE][SIZE]rune) {
	printSeparator()
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			fmt.Printf(" %c |", board[i][j])
		}
		fmt.Println()
		printSeparator()
	}
}


func DelRow(row int) bool {
	hasX := false;

	for j:=0; j<SIZE; j++{
		if board[row][j] == 'X'{

		}
	}

	for j := 0; j < SIZE; j++ {
		displayBoard[row][j] = board[row][j]
	}
	return hasX
}



func DelCol(col int) bool {
	hasX := false;

	for i:=0; i<SIZE; i++{
		if board[i][col] == 'X'{

		}
	}

	for i := 0; i < SIZE; i++ {
		displayBoard[i][col] = board[i][col]
	}
	return hasX
}