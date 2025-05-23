package main

import (
	"fmt"
)

func main() {
	players := []string{"charles", "martina", "michael", "florence", "eli"}

	// Print the first three players
	fmt.Println("Here are the first three players on my team:")
	for _, player := range players[:3] {
		fmt.Println(capitalize(player))
	}
}

// Function to capitalize the first letter of the string
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}

// Here are the first three players on my team:
// Charles
// Martina
// Michael

