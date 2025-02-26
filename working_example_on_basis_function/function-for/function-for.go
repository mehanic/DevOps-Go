package main

import (
	"fmt"
	"strings"
)

// greetUsers prints a simple greeting to each user in the list.
func greetUsers(names []string) {
	for _, name := range names {
		msg := "Hello, " + strings.Title(name) + "!"
		fmt.Println(msg)
	}
}

func main() {
	usernames := []string{"hannah", "ty", "margot"}
	greetUsers(usernames)
}

