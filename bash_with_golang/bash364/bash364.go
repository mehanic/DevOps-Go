package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <num1> <num2>")
		return
	}

	// Convert command line arguments to integers
	num1, err1 := strconv.Atoi(os.Args[1])
	num2, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil {
		fmt.Println("Both arguments must be integers.")
		return
	}

	// Perform the calculations
	result1 := num1 + num2
	result2 := num1 + num2 // Similar to declare -i in Bash

	// Print results
	fmt.Printf("%d + %d = %d  -> # let RESULT1=$1+$2\n", num1, num2, result1)
	fmt.Printf("%d + %d = %d  -> # declare -i RESULT2; RESULT2=$1+$2\n", num1, num2, result2)
	fmt.Printf("%d + %d = %d  -> # $(($1 + $2))\n", num1, num2, num1+num2)
}

