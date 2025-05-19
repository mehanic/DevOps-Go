package main

import (
	"fmt"
	"unicode"
)

// checkio checks if the string meets the specified criteria.
func checkio(data string) bool {
	if len(data) > 9 {
		hasUpper := false
		hasLower := false
		hasDigit := false

		// Check each character in the string
		for _, char := range data {
			if unicode.IsUpper(char) {
				hasUpper = true
			}
			if unicode.IsLower(char) {
				hasLower = true
			}
			if unicode.IsDigit(char) {
				hasDigit = true
			}
		}

		// Return true only if all conditions are satisfied
		return hasUpper && hasLower && hasDigit
	}
	return false
}

func main() {
	// Test cases to validate the function
	if checkio("A1213pokl") != false {
		fmt.Println("1st example failed")
	}
	if checkio("bAse730onE4") != true {
		fmt.Println("2nd example failed")
	}
	if checkio("asasasasasasasaas") != false {
		fmt.Println("3rd example failed")
	}
	if checkio("QWERTYqwerty") != false {
		fmt.Println("4th example failed")
	}
	if checkio("123456123456") != false {
		fmt.Println("5th example failed")
	}
	if checkio("QwErTy911poqqqq") != true {
		fmt.Println("6th example failed")
	}

	fmt.Println("All tests passed!")
}

