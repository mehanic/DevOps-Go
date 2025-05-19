package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Infinite loop instead of recursion
	for {
		fmt.Print("Press anything to continue loop: ")
		_, _ = reader.ReadString('\n') // Read input and ignore the value
	}
}

