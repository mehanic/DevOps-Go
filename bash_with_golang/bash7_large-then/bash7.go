package main

import (
	"fmt"
)

func main() {
	userAge := 18
	ageLimit := 18
	name := "Bob" // Change to your username if you want to execute the nested logic
	hasNightmares := true
	user := "Bob" // Change to test with different usernames

	if user == name {
		if userAge >= ageLimit {
			if hasNightmares {
				fmt.Printf("%s gets nightmares, and should not see the movie\n", user)
			}
		}
	} else {
		fmt.Println("Who is this?")
	}
}

