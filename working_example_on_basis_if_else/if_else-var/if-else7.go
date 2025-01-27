package main

import "fmt"

func main() {
	name := "Alice"  // Set the name value
	age := 10        // Set the age value

	if name == "Alice" {
		fmt.Println("Hi, Alice.")
	} else if age < 12 {
		fmt.Println("You are not Alice, kiddo.")
	} else {
		fmt.Println("You are neither Alice nor a little kid.")
	}
}

//Hi, Alice.
