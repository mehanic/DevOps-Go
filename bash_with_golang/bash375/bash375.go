package main

import (
	"fmt"
)

func main() {
	// Define string variables
	str1 := "Hello"
	str2 := "Hell"
	str3 := ""
	str4 := "Hello"

	// Print the values of str1, str2, str3, and str4
	fmt.Printf("str1 = %s , str2 = %s , str3 = %s , str4 = %s\n", str1, str2, str3, str4)

	// Check if 'str3' is empty
	fmt.Print("Is str3 empty? ")
	if str3 == "" {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}

	// Check if 'str2' is not empty
	fmt.Print("Is str2 not empty? ")
	if str2 != "" {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}

	// Check if 'str1' and 'str4' are equal
	fmt.Print("Are str1 and str4 equal? ")
	if str1 == str4 {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}

	// Check if 'str1' and 'str2' are different
	fmt.Print("Are str1 and str2 different? ")
	if str1 != str2 {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}

