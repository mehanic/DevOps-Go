package main

import (
	"fmt"
)

func main() {
	// Define variables
	NAME := "Sebastiaan"
	HOME_TYPE := "house"
	LOCATION := "Utrecht"
	PARTNER_NAME := "Sanne"
	ANIMAL_TYPES := "gecko and hamster"

	// Print the story
	fmt.Printf("%s lives in a %s in %s, together with %s and their two pets: a %s.\n", NAME, HOME_TYPE, LOCATION, PARTNER_NAME, ANIMAL_TYPES)
}

