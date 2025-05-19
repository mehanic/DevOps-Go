package main

import (
	"fmt"
)

func main() {
	// Declare variables
	animalTypes := "gecko and hamster"
	LOCATION := "Utrecht"
	name := "Sebastiaan"
	home_type := "house"
	partner_name := "Sanne"

	// Print the story
	fmt.Printf("%s lives in a %s in %s, together with %s and their two pets: a %s.\n", name, home_type, LOCATION, partner_name, animalTypes)
}

