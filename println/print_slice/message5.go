package main

import (
    "fmt"
    "strings"
)

func main() {
    // Initial slice of guests
    guests := []string{"chima", "daniel", "kosomo"}

    // Print invitation message
    fmt.Printf("\nI'd like to invite %s, %s and %s to dinner.\n",
        strings.Title(guests[0]),
        strings.Title(guests[1]),
        strings.Title(guests[2]),
    )

    // A guest will not be available
    notComing := guests[len(guests)-1]
    fmt.Printf("\tOh, no! %s will not be able to make it...\n", strings.Title(notComing))

    // Replace the guest
    newGuest := "joy"
    fmt.Printf("\tI'll call %s instead.\n", strings.Title(newGuest))

    guests[len(guests)-1] = newGuest

    // Print updated invitation messages
    fmt.Println("\nMESSAGE:")
    for _, guest := range guests {
        fmt.Printf("\tHi! %s, would you care to join me for dinner tonight?\n", strings.Title(guest))
    }
}



// I'd like to invite Chima, Daniel and Kosomo to dinner.
// 	Oh, no! Kosomo will not be able to make it...
// 	I'll call Joy instead.

// MESSAGE:
// 	Hi! Chima, would you care to join me for dinner tonight?
// 	Hi! Daniel, would you care to join me for dinner tonight?
// 	Hi! Joy, would you care to join me for dinner tonight?

