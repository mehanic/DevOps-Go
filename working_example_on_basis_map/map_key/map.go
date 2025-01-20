package main

import "fmt"

func main() {
	var people = map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}

	// Check if the key "Tom" exists in the map
	if val, ok := people["Tom"]; ok {
		// If it exists, print the value
		fmt.Println(val) // Output: 1
	}
}

// 1
