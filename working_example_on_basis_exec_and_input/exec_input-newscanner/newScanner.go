package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Creating a scanner to capture user input
	scanner := bufio.NewScanner(os.Stdin)

	// Function to get the cat name from the user
	getCatName := func(catNumber int) string {
		fmt.Printf("Enter the name of cat %d:\n", catNumber)
		scanner.Scan()
		return scanner.Text()
	}

	// Getting the names of 6 cats
	catName1 := getCatName(1)
	catName2 := getCatName(2)
	catName3 := getCatName(3)
	catName4 := getCatName(4)
	catName5 := getCatName(5)
	catName6 := getCatName(6)

	// Printing the cat names
	fmt.Println("The cat names are:")
	fmt.Println(catName1 + " " + catName2 + " " + catName3 + " " + catName4 + " " + catName5 + " " + catName6)
}


// Enter the name of cat 1:
// larry
// Enter the name of cat 2: