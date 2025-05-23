package main

import "fmt"

func main() {
	fmt.Println("If Else in GoLang")

	loginCount := 9
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login counts"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less that 10")
	}

	// if err != nil {

	// }
}


// If Else in GoLang
// Regular user
// Number is odd
// Num is less than 10
