package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	help := flag.Bool("h", false, "Display help")
	firstname := flag.String("firstname", "", "Set firstname")
	lastname := flag.String("lastname", "", "Set lastname")

	// Parse the flags
	flag.Parse()

	// Check if help is requested
	if *help {
		fmt.Println("usage: [flags]\nFlags:\n  -h\tDisplay help\n  -firstname=<value>\n  -lastname=<value>")
		os.Exit(0)
	}

	// Check if we have required arguments
	if *firstname == "" || *lastname == "" {
		fmt.Fprintln(os.Stderr, "Both firstname and lastname are required!")
		os.Exit(3)
	}

	// Print welcome message
	fmt.Printf("Welcome %s %s!\n", *firstname, *lastname)
	os.Exit(0)
}

