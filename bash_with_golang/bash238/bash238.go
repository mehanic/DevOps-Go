package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if the number of arguments is exactly 2
	if len(os.Args) != 3 {
		fmt.Println("You must provide <student> <grade>")
		os.Exit(2)
	}

	student := os.Args[1] // First argument: student name
	grade := strings.ToUpper(os.Args[2]) // Second argument: grade, converted to uppercase

	// Evaluate the grade using a switch-case structure
	switch grade {
	case "A", "B", "C":
		fmt.Printf("%s is a star pupil\n", student)
	case "D":
		fmt.Printf("%s needs to try a little harder!\n", student)
	case "E", "F":
		fmt.Printf("%s could do a lot better next year\n", student)
	default:
		fmt.Printf("Grade could not be evaluated for %s %s\n", student, grade)
	}
}

