package main

import "fmt"

func main() {
	word := "c"
	switch word {
	case "a":
		fmt.Println("A")
	case "b":
		fmt.Println("B")
	default:
		fmt.Println("error")
	}
}
