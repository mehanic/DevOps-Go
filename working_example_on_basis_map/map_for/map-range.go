package main

import (
	"fmt"
)

func main() {
	// Create an empty slice to hold aliens
	aliens := []map[string]interface{}{}

	// Generate 30 aliens and add to the slice
	for alienNumber := 0; alienNumber < 30; alienNumber++ {
		newAlien := map[string]interface{}{
			"color":  "yellow",
			"points": 5,
			"speed":  "slow",
		}
		aliens = append(aliens, newAlien)
		//fmt.Println(aliens)
	}

	// Modify the first 3 aliens based on their color
	for i := 0; i < 3; i++ {
		alien := aliens[i]
		if alien["color"] == "green" {
			alien["color"] = "yellow"
			alien["speed"] = "medium"
			alien["points"] = 10
		} else if alien["color"] == "yellow" {
			alien["color"] = "red"
			alien["speed"] = "fast"
			alien["points"] = 15
		}
	}

	// Print the first 5 aliens
	for i := 0; i < 10; i++ {
		fmt.Println(aliens[i])
	}

	fmt.Println("...")
	fmt.Printf("Total number of aliens: %d\n", len(aliens))
}

// map[color:red points:15 speed:fast]
// map[color:red points:15 speed:fast]
// map[color:red points:15 speed:fast]
// map[color:yellow points:5 speed:slow]
// map[color:yellow points:5 speed:slow]
// ...
// Total number of aliens: 30
