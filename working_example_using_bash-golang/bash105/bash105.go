package main

import (
	"fmt"
)

func main() {
	// Define the map
	myMap := map[int]string{
		1: "liyang",
		2: "liuyuan",
	}

	// Print values from the map
	fmt.Println(myMap[1])
	fmt.Println(myMap[2])

	// Print header with green background
	fmt.Print("\033[42m") // Green background
	fmt.Printf("%-10s\t%-10s\n", "key", "value")
	fmt.Print("\033[0m") // Reset formatting

	// Iterate over the map and print key-value pairs
	for key, value := range myMap {
		fmt.Printf("%-10d\t%-10s\n", key, value)
	}
}

