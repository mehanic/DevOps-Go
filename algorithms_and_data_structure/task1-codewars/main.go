// Write a function that accepts an array of 10 integers (between 0 and 9), that returns a string of those numbers in the form of a phone number.

// Example
// CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // returns "(123) 456-7890"
// The returned format must be correct in order to complete this challenge.

// Don't forget the space after the closing parentheses!

package main

import (
	"fmt"
)

func CreatePhoneNumberBase(n [10]uint) string { //uint8       the set of all unsigned  8-bit integers (0 to 255)

	// Convert slices to string and format as (xxx) xxx-xxxx
	f := fmt.Sprintf("%d%d%d", n[0], n[1], n[2])
	s := fmt.Sprintf("%d%d%d", n[3], n[4], n[5])
	t := fmt.Sprintf("%d%d%d%d", n[6], n[7], n[8], n[9])

	// Return the formatted phone number
	return fmt.Sprintf("(%s) %s-%s", f, s, t)

}

func main() {
	n := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(CreatePhoneNumberBase(n))

}
