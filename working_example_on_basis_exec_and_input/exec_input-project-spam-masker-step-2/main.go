package main

import (
	"fmt"
	"os"
)

const (
	link  = "http://"
	nlink = len(link)
	mask  = '*'
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("gimme somethin' to mask!")
		return
	}

	var (
		text = args[0]
		size = len(text)
		buf  = make([]byte, 0, size)

		in bool
	)

	for i := 0; i < size; i++ {
		// slice the input and look for the link pattern
		// do not slice it when it goes beyond the input text's capacity
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			// set the flag: we're in a link! -> "http://....."
			in = true

			// add the "http://" manually
			buf = append(buf, link...)

			// jump to the next character after "http://"
			i += nlink
		}

		// get the current byte from the input
		c := text[i]

		// disable the link detection flag
		// this will prevent masking the rest of the bytes
		switch c {
		case ' ', '\t', '\n': // try -> unicode.IsSpace
			in = false
		}

		// if we're in the link detection mode (inside the link bytes)
		// then, mask the current character
		if in {
			c = mask
		}
		buf = append(buf, c)
	}

	// print out the buffer as text (string)
	fmt.Println(string(buf))
}


//main.go "Check out this cool site http://example.com for more info!"
//Check out this cool site http://*********** for more info!
