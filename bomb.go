package main

import (
	"fmt"
	"time"
)

func printBombExplosion() {
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

func clearScreen() {
	// This will work on most Unix-like systems (Linux, macOS).
	// Windows may need a different command.
	fmt.Print("\033[H\033[2J")
}

func main() {
	// Simulate the player finding the bomb
	fmt.Println("You found the bomb!")
	time.Sleep(1 * time.Second)
	clearScreen()

	// Print bomb explosion
	printBombExplosion()
	time.Sleep(1 * time.Second)

	// Simulate end of game after explosion
	fmt.Println("Game Over!")
}
