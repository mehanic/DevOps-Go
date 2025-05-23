package main

import "fmt"

func main() {
	// Example 1: Iterating over a map
	people1 := map[string]int{
		"Tom":   1,
		"Bob":   2,
		"Sam":   4,
		"Alice": 8,
	}

	for key, value := range people1 {
		fmt.Println(key, value)
	}

	// Example 2: Adding a new key-value pair to the map
	people2 := map[string]int{
		"Tom": 1,
		"Bob": 2,
	}
	people2["Kate"] = 128
	fmt.Println(people2) // Output: map[Tom:1 Bob:2 Kate:128]

	// Example 3: Deleting a key-value pair from the map
	people3 := map[string]int{
		"Tom": 1,
		"Bob": 2,
		"Sam": 8,
	}
	delete(people3, "Bob")
	fmt.Println(people3) // Output: map[Tom:1 Sam:8]

	fmt.Println("-----------------------------")

	main1()
}

//--function

func setSettings(settings map[string]interface{}) {
	if val, ok := settings["brightness"]; ok {
		fmt.Printf("Setting brightness to %v\n", val)
	}
	// Continue setting other parameters
}

func main1() {

	settings := map[string]interface{}{
		"brightness": 75,
	}
	setSettings(settings)
}

// Tom 1
// Bob 2
// Sam 4
// Alice 8
// map[Bob:2 Kate:128 Tom:1]
// map[Sam:8 Tom:1]
// -----------------------------
// Setting brightness to 75
