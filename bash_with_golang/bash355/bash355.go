package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	var str1, str2 string

	// Read two numbers from the user
	fmt.Print("Please enter 1st Number: ")
	fmt.Scan(&num1)
	fmt.Print("Please enter 2nd Number: ")
	fmt.Scan(&num2)
	fmt.Println()

	// Test for Equal
	if num1 == num2 {
		fmt.Println("1: Numbers are equal.")
	} else {
		fmt.Println("1: Numbers are not equal.")
	}

	// Test for Not Equal
	if num1 != num2 {
		fmt.Println("2: Numbers are not equal.")
	} else {
		fmt.Println("2: Numbers are equal.")
	}

	// Test for Greater Than or Equal
	if num1 >= num2 {
		fmt.Println("3: First number is greater than or equal to the second.")
	} else {
		fmt.Println("3: First number is less than the second.")
	}

	// Read two strings from the user
	fmt.Print("Please enter 1st String: ")
	fmt.Scan(&str1)
	fmt.Print("Please enter 2nd String: ")
	fmt.Scan(&str2)

	// Test for Two Strings Are Equal
	if str1 == str2 {
		fmt.Println("4: Strings are equal.")
	} else {
		fmt.Println("4: Strings are not equal.")
	}

	// Test for The Length Of The String Is > 0
	if len(str1) == 0 {
		fmt.Println("5: First string is empty.")
	} else {
		fmt.Println("5: First string is not empty.")
	}

	// Test for The String Is Not NULL
	if str2 != "" {
		fmt.Println("6: Second string is not empty.")
	} else {
		fmt.Println("6: Second string is empty.")
	}
}

