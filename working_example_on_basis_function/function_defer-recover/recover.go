package main

import (
	"fmt"
)

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic happend FRIST:", err)
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic happend SECOND:", err)
			panic("Second panic")
		}
	}()

	fmt.Println("Some userful work")
	panic("Something bad happend")
	return
}

func main() {
	deferTest()
	return
}

// Some userful work
// Panic happend SECOND: Something bad happend
// Panic happend FRIST: Second panic
