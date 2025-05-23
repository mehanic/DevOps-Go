package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Ensure that exactly 2 command-line arguments are provided (including the script name)
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run script.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	fmt.Printf("We're going to erase %s.\n", filename)
	fmt.Println("If you don't want that, hit CTRL-C (^C).")
	fmt.Println("If you do want that, hit RETURN.")

	// Read user input to continue
	var input string
	fmt.Scanln(&input)

	fmt.Println("Opening the file...")
	file, err := os.Create(filename) // Creates and truncates the file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("Truncating the file. Goodbye!")

	fmt.Println("Now I'm going to ask you for three lines.")

	// Create a buffered writer to write to the file
	writer := bufio.NewWriter(file)

	// Read lines from user
	fmt.Print("line 1: ")
	line1, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Print("line 2: ")
	line2, _ := bufio.NewReader

}
