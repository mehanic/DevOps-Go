package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // Get command-line arguments, excluding the program name

	// Print total number of arguments and their values
	fmt.Println("$#: ", len(args))
	fmt.Println("$@: ", args)
	fmt.Println("$*: ", args)
	fmt.Println()

	// Print specific arguments
	if len(args) >= 10 {
		fmt.Println("$1 $2 $9 $10 are: ", args[0], args[1], args[8], args[9])
	} else {
		fmt.Println("$1 $2 $9 $10 are: Not enough arguments")
	}
	fmt.Println()

	// Shift: remove the first argument
	if len(args) > 0 {
		args = args[1:] // Shift by removing the first argument
	}

	// Print after shift
	fmt.Println("$#: ", len(args))
	fmt.Println("$@: ", args)
	fmt.Println("$*: ", args)
	fmt.Println()

	// Print specific arguments after first shift
	if len(args) >= 9 {
		fmt.Println("$1 $2 $9 are: ", args[0], args[1], args[8])
	} else {
		fmt.Println("$1 $2 $9 are: Not enough arguments")
	}

	// Shift by removing the first two arguments
	if len(args) > 1 {
		args = args[2:] // Shift by removing the first two arguments
	}

	// Print after second shift
	fmt.Println("$#: ", len(args))
	fmt.Println("$@: ", args)
	fmt.Println("$*: ", args)
	fmt.Println()

	// Print specific arguments after second shift
	if len(args) >= 9 {
		fmt.Println("$1 $2 $9 are: ", args[0], args[1], args[8])
	} else {
		fmt.Println("$1 $2 $9 are: Not enough arguments")
	}

	// Attempting to print ${10}
	if len(os.Args) >= 11 {
		fmt.Println("${10}: ", os.Args[10]) // os.Args[10] corresponds to the 10th argument
	} else {
		fmt.Println("${10}: Not enough arguments")
	}
}

