package main

import (
	"fmt"
	"strings"
)

func main() {
	// Name with whitespaces and newline/tabs
	name := " \n\tbenjamin tsai\n\t "
	fmt.Println("Hello " + name + ", would you like to learn some Python today?")

	// Different string manipulations
	fmt.Println(strings.Title(name))                     // Title Case
//	fmt.Println(strings.Title(strings.TrimRight(name)))  // Right Strip + Title Case
//	fmt.Println(strings.Title(strings.TrimLeft(name)))   // Left Strip + Title Case
	fmt.Println(strings.Title(strings.TrimSpace(name)))  // Full Strip + Title Case

	// Sentence and string concatenation
	sentence := "Albert Einstein once said:\n\t \"A person who never made a mistake never tried anything new.\""
	fmt.Println(sentence)

	// Famous quote
	famousPerson := "Albert Einstein"
	message := famousPerson + " once said:\n\t \"A person who never made a mistake never tried anything new.\""
	fmt.Println(message)
}

// Hello  
// 	benjamin tsai
// 	 , would you like to learn some Python today?
 
// 	Benjamin Tsai
	 
// Benjamin Tsai
// Albert Einstein once said:
// 	 "A person who never made a mistake never tried anything new."
// Albert Einstein once said:
// 	 "A person who never made a mistake never tried anything new."

