package main

import (
	"fmt"
)

func main() {
	// Defining alien_0 map
	alien0 := map[string]interface{}{
		"x_position": 0,
		"y_position": 25,
		"speed":      "medium",
	}

	// Printing original x-position
	fmt.Println("Original x-position:", alien0["x_position"])

	// Updating speed and calculating x_increment based on speed
	alien0["speed"] = "fast"
	var xIncrement int

	switch alien0["speed"] {
	case "slow":
		xIncrement = 1
	case "medium":
		xIncrement = 2
	default:
		xIncrement = 3
	}

	// Updating x_position
	alien0["x_position"] = alien0["x_position"].(int) + xIncrement
	fmt.Println("New x-position:", alien0["x_position"])

	// Creating new alien_0 map for color and points
	alien0 = map[string]interface{}{
		"color":  "green",
		"points": 5,
	}

	// Printing alien's color
	fmt.Println("The alien is", alien0["color"], ".")

	// Updating color to yellow
	alien0["color"] = "yellow"
	fmt.Println("The alien is now", alien0["color"], ".")

	// Printing full alien_0 map
	fmt.Println(alien0)

	// Printing color and points values
	fmt.Println(alien0["color"])
	fmt.Println(alien0["points"])

	// Adding new x_position and y_position values
	alien0["x_position"] = 0
	alien0["y_position"] = 25
	fmt.Println(alien0)

	// Re-initializing alien_0
	alien0 = map[string]interface{}{
		"color":  "green",
		"points": 5,
	}
	fmt.Println(alien0)

	// Deleting the points key
	delete(alien0, "points")
	fmt.Println(alien0)
}

// Original x-position: 0
// New x-position: 3
// The alien is green .
// The alien is now yellow .
// map[color:yellow points:5]
// yellow
// 5
// map[color:yellow points:5 x_position:0 y_position:25]
// map[color:green points:5]
// map[color:green]
