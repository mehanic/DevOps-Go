package main

import "fmt"


func main() {
	// for init_statement; condition; post_statement
	for i := 1; i <= 5; i++ {
		fmt.Println("No more", i)
	}
}

// No more 1
// No more 2
// No more 3
// No more 4
// No more 5
