package main

import (
	"fmt"
	"strings"
)

func main() {
	// Define the slice with names
	names := []string{"chidi", "peter", "ada", "feldor", "mickael"}

	// Print each name with different messages
	fmt.Println("Good morning " + strings.Title(names[0]) + ", how was your night?")
	fmt.Println("Hello " + strings.Title(names[1]) + ", nice T-shirt you've got.")
	fmt.Println(strings.Title(names[2]) + ", I just cracked your code.")
	fmt.Println("The name " + strings.Title(names[3]) + " does not exist in my city.")
	fmt.Println(strings.Title(names[4]) + " is a character in an unwatched movie!")
}

