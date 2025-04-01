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
			hasX = true

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
			hasX = true

		}
	}

	for i := 0; i < SIZE; i++ {
		displayBoard[i][col] = board[i][col]
	}
	return hasX
}

func PrintBombExplosion(){
	 explosion := []string{
		"    BOOM!    ",
		"    . . .    ",
		"   ( * * )   ",
		"   (  *  )   ",
		"  (   *   )  ",
		"   `-'-'`    ",
	}
	for _, line := range explosion {
		fmt.Println(line)
	}
}

func ClearScreen() {
	
	fmt.Print("\033[H\033[2J")
}


func main() {
	initializeBoard()
	initializeDisplayBoard()
	RandomSpotGenerator()
	
	attempts := 0
	gameOver := false
	
	for !gameOver {
		fmt.Println("Avoid the hidden cross:")
		printtheBoard(displayBoard)
		
		fmt.Println("\nOptions:")
		fmt.Println("1. Reveal a position")
		fmt.Println("2. Delete a row")
		fmt.Println("3. Delete a column")
		fmt.Print("Enter your choice (1-3): ")
		
		var choice int
		fmt.Scan(&choice)
		
		switch choice {
		case 1:

			var row, col int
			fmt.Print("Enter row (0-4): ")
			fmt.Scan(&row)
			fmt.Print("Enter column (0-4): ")
			fmt.Scan(&col)
			

			if row < 0 || row >= SIZE || col < 0 || col >= SIZE {
				fmt.Println("Invalid coordinates! Please enter values between 0 and 4.")
				continue
			}
			

			if displayBoard[row][col] != '?' {
				fmt.Println("You already revealed this position!")
				continue
			}
			
			attempts++
			

			if board[row][col] == 'X' {
				displayBoard[row][col] = 'X'
				gameOver = true
				fmt.Println("\nOh no! You found the cross!")
				printtheBoard(displayBoard)

				time.Sleep(1 * time.Second)
	      ClearScreen()
				PrintBombExplosion()
	      time.Sleep(1 * time.Second)
			

				fmt.Printf("Game over! You survived for %d turns.\n", attempts)
			} else {
				displayBoard[row][col] = ' '
				fmt.Println("\nGood job! No cross here. You're safe for now.")
			}
			
		case 2:

			var row int
			fmt.Print("Enter row to delete (0-4): ")
			fmt.Scan(&row)
			

			if row < 0 || row >= SIZE {
				fmt.Println("Invalid row! Please enter a value between 0 and 4.")
				continue
			}
			
			attempts++
			
			if DelRow(row) {
				gameOver = true
				fmt.Println("\nOh no! You deleted the row with the cross!")
				printtheBoard(displayBoard)
			
				time.Sleep(1 * time.Second)
	      ClearScreen()
				PrintBombExplosion()
	      time.Sleep(1 * time.Second)
				
				fmt.Printf("Game over! You survived for %d turns.\n", attempts)
			} else {
				fmt.Println("\nGood job! No cross in this row. You're safe for now.")
			}
			
		case 3:

			var col int
			fmt.Print("Enter column to delete (0-4): ")
			fmt.Scan(&col)
			

			if col < 0 || col >= SIZE {
				fmt.Println("Invalid column! Please enter a value between 0 and 4.")
				continue
			}
			
			attempts++
			
			if DelCol(col) {
				gameOver = true
				fmt.Println("\nOh no! You deleted the column with the cross!")
				printtheBoard(displayBoard)
				time.Sleep(1 * time.Second)
	      ClearScreen()
				PrintBombExplosion()
	      time.Sleep(1 * time.Second)
				
				fmt.Printf("Game over! You survived for %d turns.\n", attempts)
			} else {
				fmt.Println("\nGood job! No cross in this column. You're safe for now.")
			}
			
		default:
			fmt.Println("Invalid choice! Please enter 1, 2, or 3.")
		}
	}
}