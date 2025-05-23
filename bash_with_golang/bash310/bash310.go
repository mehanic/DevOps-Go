package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Here is an example of a Go program")

	// 1a. File listing
	fmt.Println("1a. File listing")
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fmt.Println()

	// 1b. File listing with details (long format, just the first few lines)
	fmt.Println("1b. File listing with details (long format, just the first few lines)")
	cmd = exec.Command("sh", "-c", "ls -l | head -n 5")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fmt.Println()

	// 2. Printing a calendar for the current month
	fmt.Println("2. Printing a calendar for the current month")
	cmd = exec.Command("cal")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fmt.Println()

	// 3. Here's a little for loop
	fmt.Println("3. Here's a little for loop")
	words := strings.Split("The quick brown fox jumps over the lazy dog", " ")
	for i, word := range words {
		fmt.Printf("  Word number %d is %s\n", i+1, word)
	}

	fmt.Println()
	fmt.Println("Right, I'm all done. Bye bye.")
}

