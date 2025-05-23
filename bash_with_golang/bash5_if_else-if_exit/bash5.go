package main

import (
	"fmt"
	"os"
)

func main() {
	age := 21

	if age < 18 {
		fmt.Println("You must be 18 or older to see this movie")
	} else if age == 21 {
		fmt.Println("You may see the movie and get popcorn")
	} else {
		fmt.Println("You may see the movie!")
		os.Exit(1) // Exit with status 1
	}

	fmt.Println("This line might not get executed")
}

