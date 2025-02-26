package main

import "fmt"

// display function with fixed order parameters
func display(a int, b int) {
	fmt.Printf("a=%d\n", a)
}

// Helper function to allow named-like parameters
func displayWithNamedParams(params map[string]int) {
	a := params["a"]
	b := params["b"]
	fmt.Printf("a=%d\n", a, b)
}

func main() {
	// Call display with positional arguments
	display(3, 4)

	// Call display with named-like arguments using a map
	displayWithNamedParams(map[string]int{"a": 3, "b": 4})
	displayWithNamedParams(map[string]int{"b": 4, "a": 3})
}

// a=3
// a=3
// %!(EXTRA int=4)a=3
// %!(EXTRA int=4)
