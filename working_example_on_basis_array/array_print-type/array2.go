package main

import "fmt"

// STORY:
// Hipster's Love store publishes limited books
// twice a year.
//
// The number of books they publish is fixed at 4.

// So, let's create a 4 elements string array for the books.

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	// var books [4]string
	// var books [1 + 3]string
	var books [yearly]string

	books[0] = "Kafka's Revenge"
	books[1] = "Stay Golden"
	books[2] = "Everythingship"
	books[3] += books[0] + " 2nd Edition"

	// --------------------
	// INDEXING
	// --------------------

	// Go compiler can catch indexing errors when constant is used
	// books[4] = "Neverland"

	// Go compiler cannot catch indexing errors when non-constant is used
	// i := 4
	// books[i] = "Neverland"

	// --------------------
	// PRINTING ARRAYS
	// --------------------

	// print the type
	fmt.Printf("books  : %T\n", books)

	// print the elements
	fmt.Println("books  :", books)

	// print the elements in quoted string
	fmt.Printf("books  : %q\n", books)

	// print the elements along with their types
	fmt.Printf("books  : %#v\n", books)
}

// books  : [4]string
// books  : [Kafka's Revenge Stay Golden Everythingship Kafka's Revenge 2nd Edition]
// books  : ["Kafka's Revenge" "Stay Golden" "Everythingship" "Kafka's Revenge 2nd Edition"]
// books  : [4]string{"Kafka's Revenge", "Stay Golden", "Everythingship", "Kafka's Revenge 2nd Edition"}
