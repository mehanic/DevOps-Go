package main

import (
	"fmt"
	"strings"
)

// getFormattedName returns the formatted name.
func getFormattedName(firstName, lastName, middleName string) string {
	var fullName string
	if middleName != "" {
		fullName = firstName + " " + middleName + " " + lastName
	} else {
		fullName = firstName + " " + lastName
	}
	return strings.Title(fullName)
}

func main() {
	musician := getFormattedName("jimi", "hendrix", "")
	fmt.Println(musician)

	musician = getFormattedName("john", "hooker", "lee")
	fmt.Println(musician)
}

