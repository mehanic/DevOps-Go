package main

import (
	"fmt"
	"strings"
)

// Define a struct to hold user information
type UserInfo struct {
	First    string
	Last     string
	Location string
}

func main() {
	// Create a map to hold user data
	users := map[string]UserInfo{
		"aeinstein": {
			First:    "albert",
			Last:     "einstein",
			Location: "princeton",
		},
		"mcurie": {
			First:    "marie",
			Last:     "curie",
			Location: "paris",
		},
	}

	// Iterate over the map
	for username, userInfo := range users {
		fmt.Println("\nUsername:", username)
		fullName := userInfo.First + " " + userInfo.Last
		location := userInfo.Location

		fmt.Println("\tFull name:", strings.Title(fullName))
		fmt.Println("\tLocation:", strings.Title(location))
	}
}


// Username: aeinstein
// 	Full name: Albert Einstein
// 	Location: Princeton

// Username: mcurie
// 	Full name: Marie Curie
// 	Location: Paris
