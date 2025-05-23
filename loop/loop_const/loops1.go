package main

import (
	"fmt"
	"os"
)

const (
	author    = "Mark Pilgrim (mark@diveintopython.org)"
	version   = "$Revision: 1.2 $"
	date      = "$Date: 2004/05/05 21:57:19 $"
	copyright = "Copyright (c) 2002 Mark Pilgrim"
	license   = "Python"
)

func main() {
	// Loop through command-line arguments
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}

//  /tmp/go-build4052301839/b001/exe/loops