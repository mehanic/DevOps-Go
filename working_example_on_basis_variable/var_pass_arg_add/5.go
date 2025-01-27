package main

import "fmt"

func main() {
	var znak, a, b string
	fmt.Scan(&znak) //+-/*
	fmt.Scan(&a)    //1,2,3,4
	fmt.Scan(&b)    //45,5,,6,7,3,2,

	if znak == "+" {
		fmt.Println(a + b)
	}
}

// +
// 7
// 8
// 78
