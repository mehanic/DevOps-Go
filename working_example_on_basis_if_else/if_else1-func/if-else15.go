package main

import "fmt"

// spam function in Go with special default value handling
func spam(a int, b *int) {
	if b == nil {
		fmt.Println("No b value supplied")
	} else {
		fmt.Println(a, *b)
	}
}

func main() {
	// Call spam with no second argument
	spam(1, nil) 

	// Call spam with a second argument
	b := 2
	spam(1, &b)
}

// No b value supplied
// 1 2