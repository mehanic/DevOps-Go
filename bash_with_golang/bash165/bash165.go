package main

import "fmt"

func main() {
	// Define a slice of colors
	colors := []string{"red", "green", "blue"}

	// Iterate over the slice and print each color
	for _, color := range colors {
		fmt.Printf("COLOR: %s\n", color)
	}
}

