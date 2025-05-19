package main

import (
	"fmt"
	"os"
)

func main() {
	// Print various echo statements
	fmt.Println("hello ; world")
	fmt.Println("hello   world")
	fmt.Println("escaping \\ # & \" '")
	fmt.Println("escaping \\?*\"'")
	fmt.Println("\"Hello World with strange' characters \\ * [ } ~ \\\\ . \"")
	fmt.Println("\"Hello World with strange' characters * [ } ~ \\ . \"")

	// Access and print environment variables
	fmt.Println("This is the", os.Getenv("SHELL"), "shell")
	fmt.Println("This is", os.Getenv("SHELL"), "on computer", os.Getenv("HOSTNAME"))
	fmt.Println("The userid of", os.Getenv("USER"), "is", os.Getenv("UID"))
	fmt.Println("My homedir is", os.Getenv("HOME"))

	// Color codes (for terminal output)
//	red := "\033[01;31m"
	white := "\033[01;00m"
	green := "\033[01;32m"
	blue := "\033[01;34m"

	// Print the prompt
	prompt := fmt.Sprintf("%s%s%s@%s%s%s$ ", green, os.Getenv("USER"), white, blue, os.Getenv("HOSTNAME"), white)
	fmt.Print(prompt)
}

