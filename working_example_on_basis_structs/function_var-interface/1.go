package main

import "fmt"

func main() {
	k := 8
	var s interface{}
	s = k
	word := "Hello"
	fmt.Println(s, word)
}
