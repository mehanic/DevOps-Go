package main

import (
	"fmt"
)

// deepCopy creates a deep copy of the nested map.
func deepCopy(src map[string]interface{}) map[string]interface{} {
	copy := make(map[string]interface{})
	for k, v := range src {
		switch v := v.(type) {
		case map[string]interface{}:
			// Recursive call for deep copy of nested map
			copy[k] = deepCopy(v)
		default:
			// Direct copy for other types
			copy[k] = v
		}
	}
	return copy
}

func main() {
	// Original map
	me := map[string]interface{}{
		"name": map[string]interface{}{
			"first": "Sebastian",
			"last":  "Vinci",
		},
		"age": 21,
	}

	// Clone map with deep copy
	myClone := deepCopy(me)

	// Update age in the clone
	myClone["age"] = 22

	// Print ages
	fmt.Printf("my age: %v, my clone’s age: %v\n", me["age"], myClone["age"])

	// Update first name in the clone
	nameClone := myClone["name"].(map[string]interface{})
	nameClone["first"] = "Michael"

	// Print names
	fmt.Printf("my name: %v, my clone’s name: %v\n", me["name"].(map[string]interface{})["first"], nameClone["first"])
}

// my age: 21, my clone’s age: 22
// my name: Sebastian, my clone’s name: Michael
