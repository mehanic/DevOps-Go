package main

import "fmt"

func main() {
	x, y := 15, 10

	//fmt.Printf("%T, %v\n", x, x)
	//fmt.Printf("%T, %v\n", y, y)
	//fmt.Printf("%T, %v\n", (x + y), (x + y))
	//fmt.Printf("%T, %v\n", (x - y), (x - y))
	//fmt.Printf("%T, %v\n", (x * y), (x * y))
	fmt.Printf("%T, %v\n", (9 / 3), (9 / 3))
	fmt.Printf("%T, %v\n", (9 / 3.0), (9 / 3.0))
	fmt.Printf("%T, %v\n", (x % y), (x % y))
	/*
		int, 3
		float64, 3
		int, 5*/


	//z := 5.0 / 2 // (2.5)
	//fmt.Printf("%T, %v\n", z, z)
	//z := float(5/2)    float64, 2.5

	//x := 10

	fmt.Println(x)

	x = x - 1

	fmt.Println(x)

	x--

	x += 1000

	fmt.Println(x)

	//fmt.Println(x++) , .\main.go:49:15: syntax error: unexpected ++, expecting comma or )
}

// int, 3
// float64, 3
// int, 5
// 15
// 14
// 1013
