package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Greeting message
	fmt.Println("Hello world!")

	// Asking for the name
	fmt.Println("What is your name?")
	myName, _ := reader.ReadString('\n')
	myName = strings.TrimSpace(myName) // Remove the newline character

	// Greeting the user
	fmt.Println("It is good to meet you, " + myName)

	// Print the length of the name
	fmt.Println("The length of your name is:")
	fmt.Println(len(myName))

	// Asking for the age
	fmt.Println("What is your age?")
	myAge, _ := reader.ReadString('\n')
	myAge = strings.TrimSpace(myAge) // Remove the newline character

	// Convert age from string to int and calculate the age next year
	ageInt, _ := strconv.Atoi(myAge)
	fmt.Println("You will be " + strconv.Itoa(ageInt+1) + " in a year.")
}

// Hello world!
// What is your name?
// max
// It is good to meet you, max
// The length of your name is:
// 3
// What is your age?
// 35
// You will be 36 in a year.
