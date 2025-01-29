package main

import "fmt"

func main() {

	a := [5]int{2, 10, 8, 15, 4}
	s := a[2:4] //[8,15]
	s[1] = 12
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s)) 
}
