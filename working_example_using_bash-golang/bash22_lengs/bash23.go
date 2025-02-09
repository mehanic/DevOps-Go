package main

import (
	"fmt"
//	"strings"
)

func main() {
	// Let's play with variable arrays first using Go's equivalent of substr
	str := "1234567890asdfghjkl"

	fmt.Printf("first character: %s\n", string(str[0]))
	fmt.Printf("first three characters: %s\n", str[:3])

	fmt.Printf("third character onwards: %s\n", str[3:])
	fmt.Printf("fourth to sixth character: %s\n", str[3:6])

	fmt.Printf("last character: %s\n", string(str[len(str)-1]))

	// Next, can we compare the alphabeticalness of strings?
	str2 := "abc"
	str3 := "bcd"
	str4 := "Bcd"

	if str2 < str3 {
		fmt.Println("STR2 is less than STR3")
	} else {
		fmt.Println("STR3 is greater than STR2")
	}

	// Does case have an effect? Yes, b is less than B
	if str3 < str4 {
		fmt.Println("STR3 is less than STR4")
	} else {
		fmt.Println("STR4 is greater than STR3")
	}
}

