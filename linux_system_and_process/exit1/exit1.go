package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Type exit to exit.")
		scanner.Scan()
		response := scanner.Text()

		if strings.TrimSpace(response) == "exit" {
			return // Exits the program
		}
		fmt.Println("You typed " + response + ".")
	}
}

// go run exit1.go                                                                                                                                               ──(Sat,Jan18)─┘
// Type exit to exit.
// run
// You typed run.
// Type exit to exit.
// ri
// You typed ri.
// Type exit to exit.
// exit
