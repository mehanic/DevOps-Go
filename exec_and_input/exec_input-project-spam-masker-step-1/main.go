package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("gimme somethin' to mask!")
		return
	}

	const (
		link  = "http://"
		nlink = len(link)
	)

	var (
		text = args[0]
		size = len(text)
		buf  = make([]byte, 0, size)
	)

	for i := 0; i < size; i++ {
		buf = append(buf, text[i])

		// slice the input and look for the link pattern
		// do not slice it when it goes beyond the input text's capacity
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
		}
	}

	// print out the buffer as text (string)
	fmt.Println(string(buf))
}


//Your program seems designed to mask (or remove) any text starting with the "http://" pattern in a given input string.

// go run main.go "Check this out: http://example.com and this too: http://anotherlink.com"
// Check this out: http://example.com and this too: http://anotherlink.com

