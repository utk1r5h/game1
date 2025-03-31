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