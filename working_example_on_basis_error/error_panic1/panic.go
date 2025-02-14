package main

import (
	"fmt"
	"strconv"
	// "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter4/panic"
)

func main() {
	fmt.Println("before panic")
	Catcher()
	fmt.Println("after panic")
}

// Panic panics with a divide by zero
func Panic() {
	zero, err := strconv.ParseInt("0", 10, 64)
	if err != nil {
		panic(err)
	}

	a := 1 / zero
	fmt.Println("we'll never get here", a)
}

// Catcher calls Panic
func Catcher() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occurred:", r)
		}
	}()
	Panic()
}


// before panic
// panic occurred: runtime error: integer divide by zero
// after panic
