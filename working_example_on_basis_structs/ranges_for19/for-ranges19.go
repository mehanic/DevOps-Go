package main

import (
	"fmt"
)

// printDict prints key-value pairs from a map
func printDict(dictionary map[string]interface{}) {
	for k, v := range dictionary {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func main() {
	example := map[string]interface{}{
		"name":      "Sebastian",
		"last_name": "Vinci",
		"age":       21,
	}

	printDict(example)
}

// age: 21
// name: Sebastian
// last_name: Vinci
