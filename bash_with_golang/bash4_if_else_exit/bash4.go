package main

import (
	"fmt"
	"os"
)

func main() {
	age := 9000

	if age < 18 {
		fmt.Println("You must be 18 or older to see this movie")
	} else {
		fmt.Println("You may see the movie!")
		os.Exit(1) // Exit with status 1
	}

	fmt.Println("This line will never get executed")
}

