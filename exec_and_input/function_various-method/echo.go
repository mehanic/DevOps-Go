package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	sep := "."

	fmt.Println(len(os.Args))

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(s)

	// increment/decrement
	var i int = len(os.Args)
	for i--; i > 0; i-- {
		s += sep + os.Args[i]
	}
	fmt.Println(s)

	// function call
	for concate(); 1 > 2; {
	}
}

// print the arguments
func concate() {
	var s string
	sep := "!"
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(s)
}
