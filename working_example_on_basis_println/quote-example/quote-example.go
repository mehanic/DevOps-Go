package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	// Famous person name
	famousPerson := "marshal mathers" // this is Eminem

	// Convert first letter of each word to uppercase (title case)
	famousPerson = cases.Title(language.English).String(famousPerson)

	// Quote from Eminem - Not Afraid
	quote := `"There's a game called circle and I don't know I'm way too up to back down."`

	// Concatenate strings
	message := famousPerson + " once said, " + quote

	// Print the final message
	fmt.Println(message) // print the quote by Marshal Mathers (Eminem)
}

//Marshal Mathers once said, "There's a game called circle and I don't know I'm way too up to back down."
