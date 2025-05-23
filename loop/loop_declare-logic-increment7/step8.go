package main

import (
	"fmt"
	"time"
    )

func main() {

	var number int = 0
	for number <= 11 {
		fmt.Println("-->", number)
		number++
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Next time ")
//----------------------------------------------------------------------------------------------------------------------
	fmt.Print("\n")

	var number2 int = 12
	for number2 <= 20 {
		fmt.Println("-->", number2)
		number2++
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Tugadi.")
}

// --> 0
// --> 1
// --> 2
// --> 3
// ^Csignal: interrupt
