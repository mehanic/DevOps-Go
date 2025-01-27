package main

import (
	"fmt"
)

func main() {
	//The prev array represents last year's books and holds three titles.
	prev := [3]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

	// You can't do this:
	// books = prev
    //The books array is declared with a size of four, allowing for an additional book this year.
	var books [4]string

	for i, b := range prev {
		books[i] += b + " 2nd Ed."
	}
    //Adding a New Book:
	books[3] = "Awesomeness"

	fmt.Printf("last year:\n%#v\n", prev)
	fmt.Printf("\nthis year:\n%#v\n", books)
}

// last year:
// [3]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}

// this year:
// [4]string{"Kafka's Revenge 2nd Ed.", "Stay Golden 2nd Ed.", "Everythingship 2nd Ed.", "Awesomeness"}
