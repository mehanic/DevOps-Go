package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a new reader to read from standard input
	reader := bufio.NewReader(os.Stdin)

	// Ask the question
	fmt.Print("Where do you stay? ")

	// Read the input from the user (similar to Bash's REPLY)
	reply, _ := reader.ReadString('\n')

	// Print the reply back to the user
	fmt.Printf("You stay in %s", reply)
}

