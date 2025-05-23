package main

import (
	"fmt"
	"strings"
)

func main() {
	guests := []string{"chima", "daniel", "kosomo"} // a slice of guests

	// Print initial invitation
	fmt.Printf("\nI'd like to invite %s, %s and %s to dinner.\n", strings.Title(guests[0]), strings.Title(guests[1]), strings.Title(guests[2]))

	notComing := guests[len(guests)-1] // a guest won't be coming
	fmt.Printf("\tOh, no! %s will not be able to make it...\n", strings.Title(notComing))

	newGuest := "joy" // to replace the uninvited guest
	fmt.Printf("\tI'll call %s instead.\n", strings.Title(newGuest))

	// Replace the guest who can't make it
	guests[len(guests)-1] = newGuest

	fmt.Println("\nMESSAGE:") // print string in uppercase
	fmt.Printf("\tOK! everyone, I just got a bigger dinner table, let's celebrate tonight.\n")

	// Adding guests to the list
	guests = append([]string{"jennifer"}, guests...)
	guests = append(guests[:2], append([]string{"nnamdi"}, guests[2:]...)...)
	guests = append(guests, "kim")

	// Print invitations
	for _, guest := range guests {
		fmt.Printf("\tHi! %s, would you care to join me for dinner tonight?\n", strings.Title(guest))
	}
}


// I'd like to invite Chima, Daniel and Kosomo to dinner.
// 	Oh, no! Kosomo will not be able to make it...
// 	I'll call Joy instead.

// MESSAGE:
// 	OK! everyone, I just got a bigger dinner table, let's celebrate tonight.
// 	Hi! Jennifer, would you care to join me for dinner tonight?
// 	Hi! Chima, would you care to join me for dinner tonight?
// 	Hi! Nnamdi, would you care to join me for dinner tonight?
// 	Hi! Daniel, would you care to join me for dinner tonight?
// 	Hi! Joy, would you care to join me for dinner tonight?
// 	Hi! Kim, would you care to join me for dinner tonight?
