package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Simulate $* (Dollar Star) by joining all arguments into a single string
	star := strings.Join(os.Args[1:], " ")
	fmt.Printf("Dollar Star is %s\n", star)

	// Simulate "$*" (Dollar Star in double quotes) as a single string with spaces
	fmt.Printf("Dollar Star in double quotes is \"%s\"\n", star)

	// Simulate $@ (Dollar At), printing all arguments separately
	fmt.Print("Dollar At is ")
	for _, arg := range os.Args[1:] {
		fmt.Printf("%s ", arg)
	}
	fmt.Println()

	// Simulate "$@" (Dollar At in double quotes), treating each argument as a separate quoted string
	fmt.Print("Dollar At in double quotes is ")
	for _, arg := range os.Args[1:] {
		fmt.Printf("\"%s\" ", arg)
	}
	fmt.Println()

	fmt.Println()
	// Looping through $*
	fmt.Println("Looping through Dollar Star")
	for _, arg := range strings.Split(star, " ") {
		fmt.Printf("Parameter is %s\n", arg)
	}

	fmt.Println()
	// Looping through "$*" (Single string treated as one argument)
	fmt.Println("Looping through Dollar Star with double quotes")
	fmt.Printf("Parameter is %s\n", star)

	fmt.Println()
	// Looping through $@
	fmt.Println("Looping through Dollar At")
	for _, arg := range os.Args[1:] {
		fmt.Printf("Parameter is %s\n", arg)
	}

	fmt.Println()
	// Looping through "$@" (Each parameter treated separately)
	fmt.Println("Looping through Dollar At with double quotes")
	for _, arg := range os.Args[1:] {
		fmt.Printf("Parameter is %s\n", arg)
	}
}

