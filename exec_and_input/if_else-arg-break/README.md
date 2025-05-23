package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Please enter your name.")
		scanner.Scan() // Read input from the user
		name := scanner.Text()

		if name == "Max" {
			break
		}
	}

	fmt.Println("Thank you!")
}

// Please enter your name.
// max
// Please enter your name.
// nina
// Please enter your name.
// ^Csignal: interrupt