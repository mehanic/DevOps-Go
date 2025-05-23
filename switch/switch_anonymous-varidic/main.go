package main

import (
	"fmt"
)

// Function to demonstrate switch
func demoSwitch(x int) {
	switch x {
	case 3:
		fmt.Println("x is 3")
	case 4, 5: // Executes if x is 4 or 5
		fmt.Println("x is 4 or 5")
	default:
		fmt.Println("x is unknown")
	}
}

// Example with initialization in switch
func demoSwitchWithInit() {
	switch x := someFunc(); x {
	case 3:
		fmt.Println("x is 3")
	default:
		fmt.Println("x is something else")
	}
}

// A placeholder for someFunc()
func someFunc() int {
	return 3 // Example: returning 3 for demonstration
}

// Variadic-like function (using a slice)
func sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// Anonymous function example in main
func main() {
	// Call the switch demonstration
	demoSwitch(4)
	demoSwitchWithInit()

	// Call the variadic-like function
	numbers := []int{1, 2, 3, 4}
	fmt.Println("Sum:", sum(numbers))

	// Anonymous function example
	result := func(word1, word2 string) string {
		return word1 + " " + word2
	}("hello", "world")
	fmt.Println(result)
}


// x is 4 or 5
// x is 3
// Sum: 10
// hello world
