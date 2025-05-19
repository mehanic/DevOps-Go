package main

import (
	"fmt"
)

func main() {
	str1 := "a"
	str2 := "b"

	// String comparison
	if str1 == str2 {
		fmt.Printf("%s is equal to %s!\n", str1, str2)
	}

	if str1 != str2 {
		fmt.Printf("%s is not equal to %s!\n", str1, str2)
	}

	notNull := "This is sth not nothing"
	null := ""

	// Check if notNull is non-empty
	if len(notNull) > 0 {
		fmt.Println("this is not at all notnull!")
	}

	// Check if null is empty
	if len(null) == 0 {
		fmt.Println("null/zero length!")
	}
}

