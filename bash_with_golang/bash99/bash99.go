package main

import (
	"fmt"
)

func main() {
	// Equivalent to `echo -n "liyang"` (no newline at the end)
	fmt.Print("liyang")

	// Equivalent to `echo` (just prints a newline)
	fmt.Println()

	// Equivalent to `echo -e "1\t2\t3"` (interprets \t as tab)
	fmt.Println("1\t2\t3")

	// Equivalent to `echo` (just prints a newline)
	fmt.Println()
}

