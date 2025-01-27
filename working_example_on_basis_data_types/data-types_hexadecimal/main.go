package main

import "fmt"

func main() {
	fmt.Printf("%d : %x\n", 35, 35)
	fmt.Printf("%d : %#x\n", 35, 35)
}

// 35 : 23
// 35 : 0x23
