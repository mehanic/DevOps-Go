package main

import (
	"fmt"
)

func main() {
	// Define the dictionary (map in Go)
	exampleDict := map[string]interface{}{
		"animals":       []string{"dog", "cat", "fish"},
		"a_number":      1,
		"a_name":        "Sebastian",
		"a_boolean":     true,
		"another_dict": map[string]string{
			"you could": "keep going",
			"like this": "forever",
		},
	}

	// Print the value associated with the key "animals"
	if animals, ok := exampleDict["animals"]; ok {
		fmt.Println(animals)
	} else {
		fmt.Println("Key 'animals' not found")
	}

	// Print the value associated with the key "a_number"
	if aNumber, ok := exampleDict["a_number"]; ok {
		fmt.Println(aNumber)
	} else {
		fmt.Println("Key 'a_number' not found")
	}

	// Attempt to access a key that does not exist
	if value, ok := exampleDict["this_one_does_not_exist"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Key 'this_one_does_not_exist' does not exist")
	}
}

// [dog cat fish]
// 1
// Key 'this_one_does_not_exist' does not exist
