package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}

// In this case, os.Args[0] will return the name of the executable (myprogram), or the relative path used to run it.
// os.Args[0] provides the path or name of the program itself.