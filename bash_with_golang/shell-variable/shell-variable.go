package main

import (
	"fmt"
	"os"
	"strings"
)

func countArgs(args []string) {
	fmt.Printf("%d args...\n", len(args))
}

func main() {
	// Print the script name
	scriptName := os.Args[0]
	fmt.Printf("script name is %s\n", scriptName)

	// Get the command-line arguments
	args := os.Args[1:] // Exclude the script name itself
	fmt.Printf("Total %d arguments\n", len(args))

	// Print all arguments as a single string
	fmt.Printf("$* is %s\n", strings.Join(args, " "))

	// Print each argument separately
	fmt.Printf("$@ is %v\n", args)

	// Pass args as a single argument
	fmt.Println("pass $* to function")
	countArgs(args)

	// Pass each argument separately
	fmt.Println("pass $@ to function")
	countArgs(args)
}

