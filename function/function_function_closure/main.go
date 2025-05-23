package main

import "fmt"

func main() {
	name := "Peter"
	counter := 0

	increment := func() {
		name = "Diana"
		fmt.Println(name)
		counter++
		fmt.Println(counter)
	}

	increment()
	increment()
}

