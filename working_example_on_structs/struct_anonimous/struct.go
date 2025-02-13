package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "kira",
		last:  "Patel",
		age:   21,
	}
	fmt.Println(p1)
}

//{kira Patel 21}
