package main

import (
	"fmt"
)

func main() {
	str1 := "Ganesh"
	str2 := "Mumbai"
	str3 := ""

	// Check if str1 is equal to str2
	if str1 == str2 {
		fmt.Println("Strings are equal.")
	} else {
		fmt.Println("Strings are not equal.")
	}

	// Check if str1 is not equal to str2
	if str1 != str2 {
		fmt.Println("Strings are not equal.")
	} else {
		fmt.Println("Strings are equal.")
	}

	// Check if str1 is not empty (length > 0)
	if len(str1) > 0 {
		fmt.Println("str1 is not empty.")
	} else {
		fmt.Println("str1 is empty.")
	}

	// Check if str3 is empty (length == 0)
	if len(str3) == 0 {
		fmt.Println("str3 is empty.")
	} else {
		fmt.Println("str3 is not empty.")
	}
}

