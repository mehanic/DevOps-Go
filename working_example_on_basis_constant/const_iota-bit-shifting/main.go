package main

import "fmt"

// we're doing bit shifting
// left shift 1 << (1*10) or 2^10
// 2^20 -> 1 MB, etc.
const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}

// binary		decimal
// 10000000000	1024
// 100000000000000000000	1048576
// 1000000000000000000000000000000	1073741824
// 10000000000000000000000000000000000000000	1099511627776
