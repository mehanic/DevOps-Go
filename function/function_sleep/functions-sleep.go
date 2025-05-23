package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("before Sleep()")

	time.Sleep(3 * time.Second)

	fmt.Println("waking up after Sleep()")
}

// before Sleep()
// waking up after Sleep()
