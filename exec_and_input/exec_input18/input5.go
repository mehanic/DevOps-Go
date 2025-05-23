package main

import (
	"fmt"
)

// Function to simulate `**kwargs` using a map
func display(p int, karg map[string]interface{}) {
	fmt.Println(p)
	fmt.Println(karg)
}

func main() {
	// Simulate the calls to `display` with various parameters

	// Creating maps for keyword arguments
	kargs1 := map[string]interface{}{
		"b": 5,
		"a": 4,
	}
	kargs2 := map[string]interface{}{
		"a": 4,
		"b": 5,
		"c": 6,
	}
	kargs3 := map[string]interface{}{
		"x":    5,
		"y":    "Hi",
		"z":    6.7,
		"user": "root",
	}

	// Calling the display function with various arguments
	display(4, kargs1)
	display(0, kargs2) // No positional argument equivalent to `display(a=4, b=5, c=6)`
	display(56, kargs3)
}

// 4
// map[a:4 b:5]
// 0
// map[a:4 b:5 c:6]
// 56
// map[user:root x:5 y:Hi z:6.7]
