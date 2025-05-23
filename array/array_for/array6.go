package main

import (
	"fmt"
	"strings"
)

func main() {
	// List of banned users
	bannedUsers := []string{"andrew", "carolina", "david"}

	// User to check
	user := "marie"

	// Check if the user is in the banned list
	isBanned := false
	for _, bannedUser := range bannedUsers {
		if strings.EqualFold(user, bannedUser) {
			isBanned = true
			break
		}
	}

	// Print message if the user is not banned
	if !isBanned {
		fmt.Printf("%s, you can post a response if you wish.\n", strings.Title(user))
	}
}

// Marie, you can post a response if you wish.
