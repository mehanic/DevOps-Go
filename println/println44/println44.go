package main

import (
	"fmt"
)

// Function to print the value of a key from the map or a default message if the key is not found
func printKey(dictionary map[string]interface{}, key string) {
	if value, ok := dictionary[key]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Key was not found")
	}
}

func main() {
	// Define the map 
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

	// Create a key
	exampleDict["this_one_does_not_exist"] = "it exists now" // Create the key and assign a value
	printKey(exampleDict, "this_one_does_not_exist")

	// Update a key
	exampleDict["a_boolean"] = false // Update the key with a new value
	printKey(exampleDict, "a_boolean")

	// Delete a key
	delete(exampleDict, "this_one_does_not_exist") // Delete the key
	printKey(exampleDict, "this_one_does_not_exist")
}


// it exists now
// false
// Key was not found
