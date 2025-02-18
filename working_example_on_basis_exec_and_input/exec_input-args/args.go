package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Getting input for the user")
	welcome := "Welcome to user input tutorial"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name")
	input, _ := reader.ReadString('\n')
	fmt.Println("So your name is:", input, " y/n")
	input, _ = reader.ReadString('\n')
}

// Getting input for the user
// Welcome to user input tutorial
// Please enter your name
// mark
// So your name is: mark
//   y/n
// y
