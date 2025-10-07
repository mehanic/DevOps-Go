package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	// Define the slice with names
	names := []string{"chidi", "peter", "ada", "feldor", "mickael"}

	// Create a caser for title casing in English
	title := cases.Title(language.English)

	// Print each name with different messages
	fmt.Println("Good morning " + title.String(names[0]) + ", how was your night?")
	fmt.Println("Hello " + title.String(names[1]) + ", nice T-shirt you've got.")
	fmt.Println(title.String(names[2]) + ", I just cracked your code.")
	fmt.Println("The name " + title.String(names[3]) + " does not exist in my city.")
	fmt.Println(title.String(names[4]) + " is a character in an unwatched movie!")
}
