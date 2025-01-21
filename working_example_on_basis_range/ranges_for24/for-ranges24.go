package main

import (
	"fmt"
)

// get function returns the value for the given key from the map.
// If the key does not exist, it returns the default value.
func get(dictionary map[string]interface{}, key string, defaultValue interface{}) interface{} {
	if value, exists := dictionary[key]; exists {
		return value
	}
	return defaultValue
}

func main() {
	myDict := map[string]interface{}{
		"name": "Sebastian",
		"age":  21,
	}

	// Use the get function to retrieve values from the map
	fmt.Println(get(myDict, "name", "Name was not present"))
	fmt.Println(get(myDict, "age", "Age was not present"))
	fmt.Println(get(myDict, "address", "Address was not present"))
}

// Sebastian
// 21
// Address was not present
