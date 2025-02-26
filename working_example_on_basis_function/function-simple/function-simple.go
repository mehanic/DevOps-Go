package main

import (
	"fmt"
	"strings"
)

func favoriteBook(title string) {
	// Title case conversion using strings.Title
	fmt.Println("One of my favorite books is " + strings.Title(title) + ".")
}

func main() {
	// Call the function with the book title
	favoriteBook("alice in wonderland")
}

