package main

import (
	"errors"
	"fmt"
)

// Coverage is a simple function with some brancing conditions
func Coverage(condition bool) error {
	if condition {
		return errors.New("condition was set")
	}
	return nil
}

func main() {
	// Example 1: condition is false — no error
	err := Coverage(false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("No error (condition false)")
	}

	// Example 2: condition is true — will return an error
	err = Coverage(true)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("No error (condition true)")
	}
}
