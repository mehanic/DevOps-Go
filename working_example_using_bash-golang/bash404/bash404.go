package main

import (
	"fmt"
)

func main() {
	// Print a message without a newline at the end
	fmt.Print("Currently we have seen the command \"echo\" used before")
	fmt.Println(" in the previous script") // Print with newline
	fmt.Println()

	// Print a message with tab and newlines
	fmt.Print("Can we also have \t tabs? \r\n\r\n?")
	fmt.Println(" NO, not yet!") // Print with newline
	fmt.Println()

	// Print a message with interpretation of backslash escapes
	fmt.Print("Can we also have \t tabs? \r\n\r\n?")
	fmt.Println(" YES, we can now! enable interpretation of backslash escapes")
	fmt.Println("We can also have:")

	// Print UTF-8 character (💈) using runes
	fmt.Println(string([]rune{0xF0, 0x9F, 0x92, 0x80})) // U+1F4B0, the "money bag" emoji

	// Final message
	fmt.Println("Check the man pages for more info ;)")
}

