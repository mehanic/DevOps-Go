package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func collatz(number int) int {
	if number%2 == 0 {
		number = number / 2
	} else {
		number = 3*number + 1
	}
	fmt.Println(number)
	return number
}

func collatzSequence() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a number:")

	// Get user input
	scanner.Scan()
	input := scanner.Text()

	// Convert input to an integer and handle potential errors
	num, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Error: You have entered a non-integer value. You must enter an integer value:")
		scanner.Scan()
		input = scanner.Text()
		num, err = strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Invalid input again. Exiting.")
			return
		}
	}

	temp := collatz(num)
	for temp != 1 {
		temp = collatz(temp)
	}
}

func main() {
	collatzSequence()
}


// Enter a number:
// 45
// 136
// 68
// 34
// 17
// 52
// 26
// 13
// 40
// 20
// 10
// 5
// 16
// 8
// 4
// 2
// 1
