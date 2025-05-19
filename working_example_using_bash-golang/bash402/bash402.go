package main

import (
	"fmt"
)

func main() {
	str := "1234567890asdfghjkl"

	// First character
	fmt.Print("First character: ")
	fmt.Println(firstCharacter(str))

	// First three characters
	fmt.Print("First three characters: ")
	fmt.Println(firstThreeCharacters(str))

	// Third character onwards
	fmt.Print("Third character onwards: ")
	fmt.Println(thirdCharacterOnwards(str))

	// Fourth to sixth character
	fmt.Print("Fourth to sixth character: ")
	fmt.Println(fourthToSixthCharacter(str))

	// Last character by itself
	fmt.Print("Last character by itself: ")
	fmt.Println(lastCharacter(str))

	// Remove last character only
	fmt.Print("Remove last character only: ")
	fmt.Println(removeLastCharacter(str))
}

// Extracts the first character from the string
func firstCharacter(s string) string {
	if len(s) > 0 {
		return string(s[0])
	}
	return ""
}

// Extracts the first three characters from the string
func firstThreeCharacters(s string) string {
	if len(s) > 3 {
		return s[:3]
	}
	return s
}

// Extracts the string from the third character onwards
func thirdCharacterOnwards(s string) string {
	if len(s) > 3 {
		return s[3:]
	}
	return ""
}

// Extracts the characters from the fourth to the sixth character
func fourthToSixthCharacter(s string) string {
	if len(s) >= 6 {
		return s[3:6]
	}
	return ""
}

// Extracts the last character of the string
func lastCharacter(s string) string {
	if len(s) > 0 {
		return string(s[len(s)-1])
	}
	return ""
}

// Removes the last character from the string
func removeLastCharacter(s string) string {
	if len(s) > 0 {
		return s[:len(s)-1]
	}
	return ""
}

