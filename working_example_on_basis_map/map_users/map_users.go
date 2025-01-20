package main

import (
	"fmt"
	"strings"
)

func main() {
	// Initialize the lists
	unconfirmedUsers := []string{"alice", "brian", "candace"}
	confirmedUsers := []string{}

	// While there are unconfirmed users, pop and confirm them
	for len(unconfirmedUsers) > 0 {
		// Pop the last user from the unconfirmedUsers slice
		currentUser := unconfirmedUsers[len(unconfirmedUsers)-1]
		unconfirmedUsers = unconfirmedUsers[:len(unconfirmedUsers)-1]

		// Confirm the user
		fmt.Println("Verifying user:", strings.Title(currentUser))
		confirmedUsers = append(confirmedUsers, currentUser)
	}

	// Print confirmed users
	fmt.Println("\nThe following users have been confirmed:")
	for _, confirmedUser := range confirmedUsers {
		fmt.Println(strings.Title(confirmedUser))
	}
}

// Verifying user: Candace
// Verifying user: Brian
// Verifying user: Alice

// The following users have been confirmed:
// Candace
// Brian
// Alice
