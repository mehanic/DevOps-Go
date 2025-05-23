package main

import "fmt"

// IF Example
func main() {
	number := 50
	var grade string

	if number <= 50 {
		grade = "Bad" // Bad
	}

	if number >= 51 {
		grade = "Good"
	}

	fmt.Println(grade)
	main1()
	main3()
	main4()
}

// Else Example
func main1() {
	number := 70
	if number <= 50 {
		fmt.Println("Failed!") // If number <= 50, print "Failed!"
	} else {
		fmt.Println("Good Job!") // Otherwise, print "Good Job!"
	}
}

// Else If Example
func main2() {
	number := 90
	if number <= 50 {
		fmt.Println("Failed!") // If number <= 50, print "Failed!"
	} else if number <= 70 {
		fmt.Println("Good Job!") // If number <= 70, print "Good Job!"
	} else {
		fmt.Println("Perfect!") // Otherwise, print "Perfect!"
	}
}

// Using && (AND) and || (OR) Operators
func main3() {
	number1 := 50
	if number1 <= 50 && number1 > 0 { // If number1 is between 1 and 50 inclusive
		fmt.Println("Not Passed") // Print "Not Passed"
	}

	number2 := 70
	if number2 <= 70 || number2 >= 51 { // If number2 is <= 70 or >= 51
		fmt.Println("Passed") // Print "Passed"
	}

	number3 := 99
	if number3 > 70 && number3 <= 100 { // If number3 is between 71 and 100 inclusive
		fmt.Println("Passed and won a prize") // Print "Passed and won a prize"
	}

	number4 := 10
	if number4 != 0 { // If number4 is not 0
		fmt.Println("Still have a chance") // Print "Still have a chance"
	}
}

// If Short Statement Example
func main4() {
	firstName := "Prasetiyo"

	// Simple if-else based on a string comparison
	if firstName == "Prasetiyo" {
		fmt.Println("Hello,", firstName) // If firstName == "Prasetiyo", print "Hello, Prasetiyo"
	} else if firstName == "Yosho" {
		fmt.Println("Hello Yosho") // If firstName == "Yosho", print "Hello Yosho"
	} else {
		var input string
		fmt.Println("What is your name?") // If firstName is neither, prompt the user
		fmt.Scan(&input)
		fmt.Println("- Nice to meet you", input)
	}

	// Short If Statement
	// Instead of declaring length before the if statement, do it inline.
	if length := len(firstName); length < 5 {
		fmt.Println("Your name is too short") // If length < 5, print this
	} else if length < 10 {
		fmt.Println("Your name is just right") // If length < 10, print this
	} else {
		fmt.Println("Your name is too long") // Otherwise, print this
	}
}


// Bad
// Good Job!
// Not Passed
// Passed
// Passed and won a prize
// Still have a chance
// Hello, Prasetiyo
// Your name is just right
