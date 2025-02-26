package main

import (
	"fmt"
)

// Function to build a person map with optional age
func buildPerson(firstName, lastName string, age ...int) map[string]interface{} {
	person := map[string]interface{}{
		"first": firstName,
		"last":  lastName,
	}
	
	// Check if age is provided
	if len(age) > 0 {
		person["age"] = age[0]
	}
	
	return person
}

func main() {
	// Create a musician map with optional age
	musician := buildPerson("jimi", "hendrix", 27)
	
	// Print the map
	fmt.Println(musician)
}

