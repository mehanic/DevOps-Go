package main

import (
	"fmt"
	"strings"
)

func main() {
	// Creating a string variable
	name := "david"

	// Make name start with an uppercase letter
	name = strings.Title(name)

	// String concatenation
	message := "Hello " + name + ", would you like to learn some Go today?"

	// Print the message
	fmt.Println(message)
  //Here's an alternative way to capitalize only the first letter: 
	if len(name) > 0 {
    name = strings.ToUpper(name[:1]) + name[1:]
}

}

//Hello David, would you like to learn some Go today?

