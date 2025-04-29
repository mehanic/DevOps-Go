package main

import (
	"fmt"
	"runtime"
)

func main() {
	// runtime.NumCPU() is a call expression
	fmt.Println(runtime.NumCPU() + 1)
	fmt.Println(runtime.Version())
	//fmt.Println(golang.Version())
	if 5 > 1 {
		fmt.Println("bigger")
	}
}

// 13
// go1.23.0
// bigger
