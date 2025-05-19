package main

import (
	"fmt"
//	"strings"
)

func main() {
	// Explanation of Bash flags
	fmt.Println("-e instructs bash to exit if any command has a non-zero exit status")
	fmt.Println("-u affects variables; variable must be defined")
	fmt.Println("-o pipefail prevents errors from being masked; the return code of the failing code is returned")

	fmt.Println("")
	fmt.Println("IFS=' ' Internal Field Separator controls what bash calls word splitting\n")

	// Define a slice of names
	names := []string{
		"Aaron Maxwell",
		"Wayne Gretzky",
		"David Beckham",
		"Anderson da Silva",
	}

	fmt.Println("With default IFS value...")
	// Iterate over the names and print them
	for _, name := range names {
		fmt.Println(name)
	}

	fmt.Println("")
	fmt.Println("With strict-mode IFS value...")

	// In Go, the default behavior of string splitting is similar to the default IFS
	// Here, we can manually use "\n" and "\t" as delimiters if necessary.
	// However, as we are just iterating over the slice of names, the output remains the same.
	for _, name := range names {
		fmt.Println(name)
	}
}

