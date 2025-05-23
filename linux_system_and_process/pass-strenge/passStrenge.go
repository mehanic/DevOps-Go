package main

import (
	"fmt"
	"regexp"
)

func passStrengthTest(passWord string) {
	lowerRegex := regexp.MustCompile(`[a-z]`)
	upperRegex := regexp.MustCompile(`[A-Z]`)
	numRegex := regexp.MustCompile(`[0-9]`)
	molower := lowerRegex.MatchString(passWord)
	moupper := upperRegex.MatchString(passWord)
	monum := numRegex.MatchString(passWord)
	if len(passWord) < 8 {
		fmt.Println("Your password is less than 8 characters which is too short.")
	} else if !molower {
		fmt.Println("You need at least one lower case letter.")
	} else if !moupper {
		fmt.Println("You need at least one upper case letter.")
	} else if !monum {
		fmt.Println("You need at least one number.")
	} else {
		fmt.Println("Your password is strong!")
	}
}

func main() {
	var passW string
	fmt.Println("What is your password?")
	fmt.Scanln(&passW)
	passStrengthTest(passW)
}
