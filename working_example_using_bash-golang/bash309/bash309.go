package main

import (
	"fmt"
)

// Function B
func functionB() {
	fmt.Println("Function B.")
}

// Function A with a parameter
func functionA(param string) {
	fmt.Println(param)
}

// Function D
func functionD() {
	fmt.Println("Function D.")
}

// Function C with a parameter
func functionC(param string) {
	fmt.Println(param)
}

func main() {
	// Function calls
	// Pass parameter to function A
	functionA("Function A.")
	functionB()
	// Pass parameter to function C
	functionC("Function C.")
	functionD()
}

