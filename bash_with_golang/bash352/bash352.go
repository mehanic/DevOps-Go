package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 6 {
		fmt.Println("Error: Please provide exactly 5 numbers.")
		return
	}

	sum := 0

	for i := 1; i < 6; i++ {
		number, err := strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Printf("Error: '%s' is not a valid number.\n", os.Args[i])
			return
		}
		sum += number
	}

	fmt.Printf("Sum of Five Numbers is: %d\n", sum)
}

