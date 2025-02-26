package main

import "fmt"

var declareMeAgain = 10

func nested() { // block scope starts
	var declareMeAgain = 5
	fmt.Println("inside nested:", declareMeAgain)
} // block scope ends
func main() { // block scope starts
	fmt.Println("inside main:", declareMeAgain)
	nested()
	// package-level declareMeAgain isn't effected
	// from the change inside the nested func
	fmt.Println("inside main:", declareMeAgain)
} // block scope ends

// inside main: 10
// inside nested: 5
// inside main: 10
