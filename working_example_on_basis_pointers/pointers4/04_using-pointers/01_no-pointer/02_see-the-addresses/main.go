package main

import "fmt"

func zero(z int) {
	// %p prints out the memory address, with a leading 0x
	fmt.Printf("%p\n", &z) // address in func zero
	fmt.Println(&z)        // address in func zero
	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x)        // address in main
	zero(x)
	fmt.Println(x) // x is still 5
}

// we'll see that the memory addresses printed out in main and
// zero are different (different copy!)
