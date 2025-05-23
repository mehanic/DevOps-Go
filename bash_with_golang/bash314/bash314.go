package main

import (
	"fmt"
	"time"
)

func main() {
	// Declare Go string variable
	BASH_VAR := "Bash Script"

	// Print the variable BASH_VAR
	fmt.Println(BASH_VAR)

	// Print using string interpolation
	currentTime := time.Now().Format(time.RFC1123)
	fmt.Printf("It's %s and \"%s\" using backticks: %s\n", BASH_VAR, BASH_VAR, currentTime)
}

