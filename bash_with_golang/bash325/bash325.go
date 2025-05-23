package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define the flags. These will act like options `-x` and `-y`.
	var xFlag bool
	var yFlag string

	// Set the flags with descriptions.
	flag.BoolVar(&xFlag, "x", false, "Option x was called.")
	flag.StringVar(&yFlag, "y", "", "Option y was called with an argument.")

	// Parse the command-line flags.
	flag.Parse()

	// Check if option x was used
	if xFlag {
		fmt.Println("Option x was called.")
	}

	// Check if option y was used and display the argument
	if yFlag != "" {
		fmt.Printf("Option y was called. Argument called is %s\n", yFlag)
	}

	// Handle the case when an invalid flag is provided
	if len(os.Args) > 1 {
		invalidArgs := flag.Args()
		if len(invalidArgs) > 0 {
			fmt.Printf("%s is not a valid option.\n", invalidArgs[0])
			fmt.Println("Usage: -x -y [argument]")
		}
	}
}

