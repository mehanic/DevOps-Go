package main

import (
	"fmt"
)

func main() {
	var catName1, catName2, catName3, catName4, catName5, catName6 string

	// Prompt for each cat name and read input
	fmt.Print("Enter the name of cat 1: ")
	fmt.Scanln(&catName1)

	fmt.Print("Enter the name of cat 2: ")
	fmt.Scanln(&catName2)

	fmt.Print("Enter the name of cat 3: ")
	fmt.Scanln(&catName3)

	fmt.Print("Enter the name of cat 4: ")
	fmt.Scanln(&catName4)

	fmt.Print("Enter the name of cat 5: ")
	fmt.Scanln(&catName5)

	fmt.Print("Enter the name of cat 6: ")
	fmt.Scanln(&catName6)

	// Display the cat names
	fmt.Println("The cat names are:")
	fmt.Println(catName1, catName2, catName3, catName4, catName5, catName6)
}

