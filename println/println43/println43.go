package main

import (
	"fmt"
)

func main() {
	// Define the map (dictionary in Go)
	exampleDict := map[string]interface{}{
		"animals":    []string{"dog", "cat", "fish"},
		"a_number":   1,
		"a_name":     "Sebastian",
		"a_boolean":  true,
		"another_dict": map[string]string{
			"you could": "keep going",
			"like this": "forever",
		},
	}

	// Define the default message
	defaultMessage := "oops, key not found"

	// Print the value associated with the key "animals"
	if animals, ok := exampleDict["animals"]; ok {
		fmt.Println(animals)
	} else {
		fmt.Println(defaultMessage)
	}

	// Print the value associated with the key "a_number"
	if aNumber, ok := exampleDict["a_number"]; ok {
		fmt.Println(aNumber)
	} else {
		fmt.Println(defaultMessage)
	}

	// Attempt to access a key that does not exist
	if value, ok := exampleDict["this_one_does_not_exist"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("KeyError")
	}

	// Get value with default message for existing key "animals"
	if value, ok := exampleDict["animals"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println(defaultMessage)
	}

	// Get value with default message for existing key "a_number"
	if value, ok := exampleDict["a_number"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println(defaultMessage)
	}

	// Get value with default message for non-existing key
	if value, ok := exampleDict["this_one_does_not_exist"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println(defaultMessage)
	}
}

// [dog cat fish]
// 1
// KeyError
// [dog cat fish]
// 1
// oops, key not found
