package main

import "fmt"

func main() {
	// go does not have the notion or while loops, we can create them using for loops
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
