package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Help message
	helpStr := "usage: [--firstname value] [--lastname value] [-h] [--help]"

	// Define command-line flags
	firstname := flag.String("firstname", "", "First name")
	lastname := flag.String("lastname", "", "Last name")
	help := flag.Bool("help", false, "Display help information")

	// Parse command-line flags
	flag.Parse()

	// Check for help flag
	if *help {
		fmt.Println(helpStr)
		os.Exit(0)
	}

	// Do we have even one argument?
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, helpStr)
		os.Exit(2)
	}

	// Sanity check for both Firstname and Lastname
	if *firstname == "" || *lastname == "" {
		fmt.Fprintln(os.Stderr, "Both firstname and lastname are required!")
		os.Exit(3)
	}

	// Print the welcome message
	fmt.Printf("Welcome %s %s!\n", *firstname, *lastname)

	os.Exit(0)
}

