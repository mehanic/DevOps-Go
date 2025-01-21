package main

import (
	"fmt"
)

func main() {
	// First loop: prints odd numbers less than 10
	currentNumber := 0
	for currentNumber < 10 {
		currentNumber++
		if currentNumber%2 == 0 {
			continue
		}
		fmt.Println(currentNumber)
	}

}

// 1
// 3
// 5
// 7
// 9
